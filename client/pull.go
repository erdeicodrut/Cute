package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"github.com/abiosoft/ishell"
)

// Get the file you specify from the server
func pull(c *ishell.Context) {
	if len(c.Args) == 0 {
		fmt.Println("Usage: pull FILE [FILE]...")
		return
	}
	name := c.Args[0]

	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return
	}
	defer conn.Close()

	toSend := Message{
		Interaction: "pull",
		Name:        name,
		Data:        []byte{},
	}

	json.NewEncoder(conn).Encode(toSend)

	for {

		var message Message
		err := json.NewDecoder(conn).Decode(&message)
		if err != nil {
			fmt.Println(err)
		}

		if message.Interaction == "Done" {
			fmt.Println("Done")
			break
		}

		if slash := strings.LastIndex(message.Name, "/"); slash > 0 {
			err := os.MkdirAll(message.Name[:slash], 0755)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		file, err := os.Create(message.Name)
		if err != nil {
			fmt.Printf("Couldn't create '%v' file because '%v'\n", message.Name, err)
			file, err = os.Open(message.Name)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		file.Write(message.Data)
		file.Close()

		fmt.Printf("Pulled file '%v'\n", message.Name)
		json.NewEncoder(conn).Encode(Message{Interaction: "thx"})

	}
}
