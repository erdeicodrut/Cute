package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"os"
)

type Config struct {
	IP   string `json:"ip"`
	PORT string `json:"port"`
}

type Message struct {
	Interaction string `json:"interaction"`
	Name        string `json:"name"`
	Data        []byte `json:"data"`
}

var configData = Config{"", ""}

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
			Action: check,
		},

		{
			Name:   "ls",
			Usage:  "Lists all the files in the server",
			Action: ls,
		},
	}

	app.Run(os.Args)
}

type ByName []os.FileInfo

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }
