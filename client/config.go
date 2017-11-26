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

	IP = strings.Trim(IP, "\n"+" ")
	PORT = strings.Trim(PORT, "\n"+" ")

	configData = Config{IP, PORT}

	json.NewEncoder(file).Encode(configData)
}
