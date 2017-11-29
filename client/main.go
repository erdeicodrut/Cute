package main

import (
	"encoding/json"
	"os"
	"github.com/abiosoft/ishell"
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
}

var configData = Config{"", ""}

func main() {
	shell := ishell.New()
	shell.Println("Welcome to Cute")

	configFile, err := os.Open("config.json")
	if err != nil {
		config(nil)
	}

	json.NewDecoder(configFile).Decode(&configData)
	configFile.Close()

	shell.AddCmd(&ishell.Cmd{
		Name: "config",
		Help: "configure Cute",
		Func: config,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "lists all the files on the server",
		Func: ls,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "push",
		Help: "pushes files to the server",
		Func: push,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "pull",
		Help: "pulls files from the server",
		Func: pull,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "check",
		Help: "checks a file and tells you if you have the latest version",
		Func: check,
	})

	shell.Run()
}
