package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Sends the a file through the connection
// Hopefully the file the client has required
func send(message Message, conn net.Conn) {
	file, err := os.Open(configData.STORAGE_PATH + message.Name)
	if err != nil {
		errorMessage := fmt.Sprintf("File '%v' doesn't exist.\n", message.Name)
		json.NewEncoder(conn).Encode(Message{Error: errorMessage})
		return
	}
	file.Close()

	bytes, _ := ioutil.ReadFile(configData.STORAGE_PATH + message.Name)

	json.NewEncoder(conn).Encode(Message{
		Name: message.Name,
		Data: bytes,
	})

	fmt.Printf("Sent file '%v'\n", message.Name)
}
