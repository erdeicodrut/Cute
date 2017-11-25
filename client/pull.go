package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"net"
	"os"
	"time"
)

func pull(c *cli.Context) {

	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
	}
	defer conn.Close()

	name := c.Args()[0]

	toSend := Message{
		Interaction: "pull",
		Name:        name,
		Data:        []byte{},
	}

	json.NewEncoder(conn).Encode(toSend)

	time.Sleep(time.Second / 2)

	exists, err := check(name) // see how
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Fprintf(os.Stdout, "File %v doesn't exist\n", name)
		return
	}

	receivedData := bufio.ReadWriter{}

	for {

		count, err := io.Copy(receivedData, conn)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading from conn %v\n", err)
			break
		}
		if count == 0 {
			fmt.Fprintln(os.Stdout, "Read all of the bytes")
			break
		}

		fmt.Printf("Read %v receivedData\n", count)

	}

	data := Message{}

	json.NewDecoder(receivedData).Decode(&data)

	exists = checkFile(data.Name)

	var file *os.File
	if !exists {
		file, _ = os.Create(data.Name)
	} else {
		file, _ = os.Open(data.Name)
	}
	defer file.Close()

	io.Copy(file, bytes.NewBuffer(data.Data))

	fmt.Fprintf(os.Stdout, "File %v pulled", file.Name())
}

func checkFile(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	file.Close()
	return true
}
