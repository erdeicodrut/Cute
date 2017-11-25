package main

import (
	"os"
	"fmt"
)

func push(message Message) {
	file, err:=os.Create(configData.STORAGE_PATH + message.Name)
	if err != nil {
		fmt.Printf("File %v couldn't be created because %v", message.Name, err.Error())
		return
	}
	file.Write(message.Data)
}