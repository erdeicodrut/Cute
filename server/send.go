package main

import (
	"net"
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)

func send(message Message, conn net.Conn) {
	file, err := os.Open(configData.STORAGE_PATH + message.Name)
	if err != nil {
		errorMessage := fmt.Sprintf("File '%v' doesn't exist.\n", message.Name)
		fmt.Print(errorMessage)
		json.NewEncoder(conn).Encode(Message{
			Error: errorMessage})
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
