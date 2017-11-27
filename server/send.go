package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Sends the a file through the connection
// Hopefully the file the client has required
func send(message Message, conn net.Conn) {
	file, err := os.Open(configData.STORAGE_PATH + message.Name)
	if err != nil {
		errorMessage := fmt.Sprintf("File '%v' doesn't exist.\n", message.Name)
		json.NewEncoder(conn).Encode(Message{Error: errorMessage})
		return
	}
	file.Close()

	pushAll(configData.STORAGE_PATH+message.Name, conn)

	json.NewEncoder(conn).Encode(Message{Interaction: "Done"})

	fmt.Printf("Sent file '%v'\n", message.Name)
}

func pushF(file *os.File, conn net.Conn) {

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

func pushAll(filename string, conn net.Conn) {
	fileTemp, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileTemp.Close()
	file, _ := fileTemp.Stat()

	if !file.IsDir() {
		pushF(fileTemp, conn)
		return
	}

	files, _ := fileTemp.Readdir(-1)

	for _, tempFile := range files {
		fmt.Println(tempFile.Name())
		pushAll(filename+"/"+tempFile.Name(), conn)
	}

}
