package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

// ls comes form list, the ls unix command, equivalent of dir in windows cmd
func ls(_ *cli.Context) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}
	defer conn.Close()

	toSend := Message{
		Interaction: "ls",
	}

	json.NewEncoder(conn).Encode(toSend)

	var message Message
	json.NewDecoder(conn).Decode(&message)

	if message.Error != "" {
		fmt.Print(message.Error)
		return
	}

	fmt.Print(string(message.Data))
}
