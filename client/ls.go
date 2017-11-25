package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"net"
	"encoding/json"
)

func ls(c *cli.Context) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}
	defer conn.Close()

	toSend := Message{
		Interaction: "ls",
	}

	json.NewEncoder(conn).Encode(toSend)

	////////////

	var message Message
	json.NewDecoder(conn).Decode(&message)

	if message.Error != "" {
		fmt.Print(message.Error)
		return
	}

	fmt.Print(string(message.Data))
}
