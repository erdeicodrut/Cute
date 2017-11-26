package main

import (
	"fmt"
	"os"
	"strings"
)

func receive(message Message) {
	os.MkdirAll(configData.STORAGE_PATH, 0755)

	for message.Name[:2] == ".." {
		message.Name = strings.Replace(message.Name, "../", "", -1)
	}

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

func createDir(s string, from int) {
	firstSlash := from + strings.Index(s[from:], "/")
	if firstSlash < 0 {
		return
	}
	err := os.MkdirAll(s[:firstSlash], 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	createDir(s, firstSlash+1)

}
