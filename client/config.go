package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/abiosoft/ishell"
)

// Configure the connection
// The configuration is stored in a json file which will be accessed at every run of the server
func config(c *ishell.Context) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("IP: ")
	IP, _ := reader.ReadString('\n')

	fmt.Print("PORT: ")
	PORT, _ := reader.ReadString('\n')

	IP = strings.Trim(IP, "\n"+" ")
	PORT = strings.Trim(PORT, "\n"+" ")

	configData = Config{IP, PORT}

	file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("Couldn't create/open 'config.json' file because '%v'\n", err)
	}
	defer file.Close()

	configJsonBytes, _ := json.Marshal(configData)
	file.Write(configJsonBytes)

	fmt.Println("Configuration saved.")
}
