package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net"
	"os"
)

// Checks if the files have the same hash in order to determine if they are the same
func checkIT(c *cli.Context) {
	same, err := check(c.Args()[0])
	if err != nil {
		panic(err)
	}

	if same {
		fmt.Fprintf(os.Stdout, "There is no change needed to be made to %v", c.Args()[0])
		return
	}

	fmt.Fprintf(os.Stdout, "There is a newer version of %v.\n Run Cute pull \"%v\" if you wish to update the file.\n", c.Args()[0], c.Args()[0])
}

// this is the actual function that does the checking
func check(fileName string) (bool, error) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return false, err
	}
	defer conn.Close()

	localFile, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to open %v because %v", fileName, err)
		return false, err

	}

	data := Message{
		Name:        fileName,
		Interaction: "check",
	}

	json.NewEncoder(conn).Encode(data)

	bytesLocalFile, _ := ioutil.ReadAll(localFile)
	md5Local := md5.Sum(bytesLocalFile)

	json.NewDecoder(conn).Decode(&data)

	if bytes.Equal(data.Data, md5Local[:]) {
		return true, nil
	}
	return false, nil
}
