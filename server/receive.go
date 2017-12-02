package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func receive(message Message) {
	os.MkdirAll(configData.STORAGE_PATH, 0755)

	message.Name = strings.Replace(message.Name, "../", "", -1)

	if x := strings.LastIndex(message.Name, "/"); x > 0 {
		os.MkdirAll(configData.STORAGE_PATH+message.Name[:x], 0755)
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

	db := <-dbC

	rows, err := db.Query("SELECT COUNT(ID) FROM File WHERE Name = ?", message.Name)
	if err != nil {
		fmt.Println(err)
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count > 0 {
		_, err = db.Exec("UPDATE File SET Date = ? Where Name = ?", time.Now().Unix(), message.Name)
		if err != nil {
			fmt.Println(err)
		}
	}

	if count == 0 {
		_, err = db.Exec("INSERT INTO File (ID, Name, Date) VALUES(NULL, ?, ?)", message.Name, time.Now().Unix())
		if err != nil {
			fmt.Println(err)
		}
	}

	dbC <- db

	fmt.Printf("Received file '%v'\n", message.Name)
}
