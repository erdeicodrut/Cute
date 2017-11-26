package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net"
	"os"
)

// Sends the file to the server
// The server will hopefully store it
func push(c *cli.Context) {
	pushF(c.Args()[0])
}

func pushF(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stdout, "No such file as %v", err)
		return
	}
	defer file.Close()

	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}

	fileBytes, err := ioutil.ReadAll(file)

	toSend := Message{
		Interaction: "push",
		Name:        file.Name(),
		Data:        fileBytes,
	}

	json.NewEncoder(conn).Encode(toSend)

	fmt.Printf("Pushed file '%v'\n", file.Name())
}
