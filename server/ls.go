package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

var filesAll []string

func ls(conn net.Conn) {

	getFiles("./")

	json.NewEncoder(conn).Encode(filesAll)

}

func getFiles(filename string) {
	fileTemp, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileTemp.Close()
	file, _ := fileTemp.Stat()

	if !file.IsDir() {
		filesAll = append(filesAll, fileTemp.Name())
		return
	}

	files, _ := fileTemp.Readdir(-1)

	for _, tempFile := range files {
		fmt.Println(tempFile.Name())
		getFiles(filename + "/" + tempFile.Name())
	}

}
