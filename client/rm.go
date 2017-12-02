package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"github.com/urfave/cli"
)

func rm(c *cli.Context) {
	name := c.Args()[0]
	rmFile(name)
}

func rmFile(name string) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return
	}
	defer conn.Close()

	data := Message{
		Name:        name,
		Interaction: "rm",
	}

	json.NewEncoder(conn).Encode(data)
}