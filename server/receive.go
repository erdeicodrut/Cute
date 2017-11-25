package main

import (
	"os"
	"fmt"
)

func receive(message Message) {
	os.MkdirAll(configData.STORAGE_PATH, 0755)

	file, err := os.Create(configData.STORAGE_PATH + message.Name)
	if err != nil {
		fmt.Printf("File %v couldn't be created because %v", message.Name, err.Error())
		return
	}
	file.Write(message.Data)

	fmt.Printf("Received file '%v'\n", message.Name)
}