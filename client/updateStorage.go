package main

import (
	"net"
	"fmt"
	"os"
)

var filesAll []string

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

func updateStorage() {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return
	}
	defer conn.Close()

	lsFiles := lsFiles(conn)

	for _, file := range lsFiles {
		same, err := check(file)
		if err != nil {
			// Will be resolved on difference
		}

		if !same {
			pullFile(conn, file)
		}
	}

	getFiles(configData.STORAGE_PATH)

	fmt.Printf("Local ls: %v", lsFiles)
	fmt.Printf("Remote ls: %v", filesAll)
}
