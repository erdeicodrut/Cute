package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

func pull(c *cli.Context) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}
	defer conn.Close()

	name := c.Args()[0]

	toSend := Message{
		Interaction: "pull",
		Name:        name,
		Data:        []byte{},
	}

	json.NewEncoder(conn).Encode(toSend)

	////////////

	var message Message
	json.NewDecoder(conn).Decode(&message)

	if message.Error != "" {
		fmt.Print(message.Error)
		return
	}

	file, err := os.Create(message.Name)
	if err != nil {
		fmt.Printf("Couldn't create '%v' file because '%v'\n", message.Name, err)
	}
	defer file.Close()

	file.Write(message.Data)

	fmt.Printf("Pulled file '%v'\n", message.Name)
}