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

	conn, err := net.Dial("tcp", initData.IP+":"+initData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", initData.IP, initData.PORT, err)
	}

	name := c.Args()[0]

	toSend := Message{
		Interaction: "pull",
		Name:        name,
		Data:        []byte{},
	}

	json.NewEncoder(conn).Encode(toSend)

	time.Sleep(time.Second / 2)

	exists := check(name) // see how

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

	exists = checkFile(data.Name) // see how

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
	return true
}
