package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

// Configure the connection
// The configuration is stored in a json file which will be accessed at every run of the server
func config(_ *cli.Context) {
	file, err := os.Open("init.json")
	if err != nil {
		file, err = os.Create("init.json")
		if err != nil {
			fmt.Printf("Couldn't create 'init.json' file because %v\n", err)
		}
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("IP: ")
	IP, _ := reader.ReadString('\n')

	fmt.Print("PORT: ")
	PORT, _ := reader.ReadString('\n')

	fmt.Print("STORAGE_PATH: ")
	STORAGE_PATH, _ := reader.ReadString('\n')

	IP = strings.Trim(IP, "\n"+" ")
	PORT = strings.Trim(PORT, "\n"+" ")
	STORAGE_PATH = strings.Trim(STORAGE_PATH, "\n"+" ")

	// Append '/' at the end of path if there isn't one
	if STORAGE_PATH[len(STORAGE_PATH)-1] != '/' {
		STORAGE_PATH += "/"
	}

	// Create the directories
	os.MkdirAll(STORAGE_PATH, 0755)

	configData = Config{IP, PORT, STORAGE_PATH}

	json.NewEncoder(file).Encode(configData)
}
