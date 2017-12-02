package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
)

func ls(_ *cli.Context) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}
	defer conn.Close()

	fmt.Println(lsFiles(conn))
}

func lsFiles(conn net.Conn) []string {
	toSend := Message{
		Interaction: "ls",
	}

	json.NewEncoder(conn).Encode(toSend)

	var fileArray []string
	json.NewDecoder(conn).Decode(&fileArray)

	return fileArray
}