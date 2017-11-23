package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

func pull(c *cli.Context) error {

	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to "+configData.IP+":"+configData.PORT+" because %v", err)
	}

	name := c.Args()[0]

	toSend := Message{
		Interaction: "pull",
		Name:        name,
		Data:        []byte{},
	}

	json.NewEncoder(conn).Encode(toSend)

	return nil
}
