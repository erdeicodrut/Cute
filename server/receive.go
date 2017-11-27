package main

import (
	"fmt"
	"os"
	"strings"
)

func receive(message Message) {
	os.MkdirAll(configData.STORAGE_PATH, 0755)

	message.Name = strings.Replace(message.Name, "../", "", -1)

	fmt.Println(message.Name)

	os.MkdirAll(configData.STORAGE_PATH+message.Name[:strings.LastIndex(message.Name, "/")], 0755)

	file, err := os.Create(configData.STORAGE_PATH + message.Name)
	if err != nil {
		fmt.Println(err)
		file, err = os.Open(configData.STORAGE_PATH + message.Name)
		if err != nil {
			return
		}
	}
	defer file.Close()

	file.Write(message.Data)

	fmt.Printf("Received file '%v'\n", message.Name)
}
