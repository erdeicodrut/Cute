package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"github.com/abiosoft/ishell"
)

// Sends the file to the server
func push(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Println("Usage: push FILE [FILE]...")
		return
	}
	pushAll(c.Args[0])
}

func pushAll(filename string) {
	fileTemp, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileTemp.Close()
	file, _ := fileTemp.Stat()

	if !file.IsDir() {
		pushFile(fileTemp)
		return
	}

	files, _ := fileTemp.Readdir(-1)

	for _, tempFile := range files {
		fmt.Println(tempFile.Name())
		pushAll(filename + "/" + tempFile.Name())
	}

}

func pushFile(file *os.File) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return
	}
	defer conn.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(conn, "Failed to read file %v, because %v\n", file.Name(), err)
	}

	toSend := Message{
		Interaction: "push",
		Name:        file.Name(),
		Data:        fileBytes,
	}

	json.NewEncoder(conn).Encode(toSend)

	fmt.Printf("Pushed file '%v'\n", file.Name())
}
