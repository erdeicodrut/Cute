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
	Error       string `json:"error"`
	Date        string `json:"date"`
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
			Action: checkIT,
		},

		{
			Name:   "ls",
			Usage:  "Lists all the files in the server",
			Action: ls,
		},
	}

	app.Run(os.Args)
}
