package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"net"
	"os"
	"strings"
)

// Get the file you specify from the server
func pull(c *cli.Context) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return
	}
	defer conn.Close()

	name := c.Args()[0]

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
