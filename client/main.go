package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"os"
	"fmt"
	"github.com/andreaskoch/go-fswatch"
	"strings"
	"path/filepath"
)

type Config struct {
	IP           string `json:"ip"`
	PORT         string `json:"port"`
	STORAGE_PATH string `json:"storage_path"`
}

type Message struct {
	Interaction string `json:"interaction"`
	Name        string `json:"name"`
	Data        []byte `json:"data"`
	Error       string `json:"error"`
	Date        uint64 `json:"date"`
}

var configData Config

func main() {
	app := cli.NewApp()

	initFile, err := os.Open("init.json")
	if err != nil {
		config(nil)
	}

	json.NewDecoder(initFile).Decode(&configData)
	initFile.Close()

	app.Name = "Cute"
	app.Usage = "A simple cloud storage kind of stuff"
	app.HideVersion = true

	app.Commands = []cli.Command{
		{
			Name:   "config",
			Usage:  "Configure Cute",
			Action: config,
		},

		{
			Name:   "push",
			Usage:  "Pushes a file to the server",
			Action: push,
		},

		{
			Name:   "pull",
			Usage:  "Pulls a file from the server",
			Action: pull,
		},

		{
			Name:   "check",
			Usage:  "Checks a file and tells you if you have the latest version",
			Action: checkIT,
		},

		{
			Name:   "ls",
			Usage:  "Lists all the files in the server",
			Action: ls,
		},

		{
			Name:   "rm",
			Usage:  "Removes a file on the server",
			Action: rm,
		},
	}

	app.Action = func(c *cli.Context) error {
		skipDotFilesAndFolders := func(path string) bool {
			return strings.HasPrefix(filepath.Base(path), ".")
		}

		folderWatcher := fswatch.NewFolderWatcher(
			configData.STORAGE_PATH,
			true,
			skipDotFilesAndFolders,
			2,
		)

		folderWatcher.Start()

		for folderWatcher.IsRunning() {
			updateStorage()

			select {

			case changes := <-folderWatcher.ChangeDetails():

				if len(changes.Moved()) != 0 {
					for _, movedFile := range changes.Moved() {
						rmFile(movedFile)
					}
				}

				if len(changes.New()) != 0 {
					for _, newFile := range changes.New() {
						pushAll(newFile)
					}
				}

				if len(changes.Modified()) != 0 {
					for _, modifiedFile := range changes.Modified() {
						pushAll(modifiedFile)
					}
				}

				//fmt.Printf("%s\n", changes.String())
				fmt.Printf("Moved: %#v\n", changes.Moved())
				fmt.Printf("New: %#v\n", changes.New())
				fmt.Printf("Modified: %#v\n\n", changes.Modified())

			}
		}

		return nil
	}

	app.Run(os.Args)
}
