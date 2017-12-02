package main

import (
	"os"
	"fmt"
)

func rm(message Message) {
	err := os.Remove(configData.STORAGE_PATH + message.Name)
	if err != nil {
		fmt.Printf("Failed to remove '%v' because '%v'\n", message.Name, err)
	}
}