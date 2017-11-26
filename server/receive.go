package main

import (
	"fmt"
	"os"
	"strings"
)

func receive(message Message) {
	os.MkdirAll(configData.STORAGE_PATH, 0755)

	if message.Name[:2] == ".." {
		message.Name = strings.Replace(message.Name, "../", "", -1)

		createDir(message.Name)

		message.Name = message.Name[strings.LastIndex(message.Name, "/")+1:]
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

func createDir(s string) {
	firstSlash := strings.Index(s, "/")
	if firstSlash < 0 {
		return
	}
	fmt.Println(s, "\nFirst slash", firstSlash)
	err := os.Mkdir(s[:firstSlash], 0755)
	if err != nil {
		os.Exit(1)
		fmt.Println(err)
	}
	createDir(s[firstSlash+1:])

}
