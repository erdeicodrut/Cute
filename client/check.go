package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func checkIT(c *cli.Context) {
	same, err := check(c.Args()[0])
	if err != nil {
		panic(err)
	}

	if same {
		fmt.Fprintf(os.Stdout, "There is no change needed to be made to %v", c.Args()[0])
		return
	}

	fmt.Fprintf(os.Stdout, "There is a newer version of %v.\n Run Cute pull \"%v\" if you wish to update the file.", c.Args()[0])
}

func check(s string) (bool, error) {
	conn, err := net.Dial("tcp", configData.IP+":"+configData.PORT)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to connect to %v:%v because %v", configData.IP, configData.PORT, err)
		return false, err
	}
	defer conn.Close()

	localFile, err := os.Open(s)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to open %v because %v", s, err)
		return false, err

	}

	data := Message{
		Name:        s,
		Data:        []byte{},
		Interaction: "check",
	}

	json.NewEncoder(conn).Encode(data)

	bytesLocalFile, _ := ioutil.ReadAll(localFile)

	md5Local := md5.Sum(bytesLocalFile)

	time.Sleep(time.Second / 2)

	json.NewDecoder(conn).Decode(&data)

	md5Remote := md5.Sum(data.Data)

	if equals(md5Remote, md5Local) {
		return true, nil
	}
	return false, nil
}

func equals(h [16]byte, o [16]byte) bool {

	for i := range h {
		if h[i] != o[i] {
			return false
		}
	}
	return true
}
