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
	defer file.Close()

	pushAll(configData.STORAGE_PATH+message.Name, conn)

	json.NewEncoder(conn).Encode(Message{Interaction: "Done"})
	fmt.Println("DONE")

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

	fmt.Printf("Struct: %v", toSend.Name)

	json.NewEncoder(conn).Encode(toSend)

	var thx Message
	err = json.NewDecoder(conn).Decode(&thx)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if thx.Interaction == "thx" {
		fmt.Println("THX")
	}

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
