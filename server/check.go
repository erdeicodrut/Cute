package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Sends the hash of the file to the client for it to decide if it has the same version of the file
func check(s string, conn net.Conn) {

	file, err := os.Open(s)
	if err != nil {
		// There is no point in continuing or logging the error.
		// A better idea would be if the client knew what happend and it can handle it properly
		// Which usually means logging it...
		answer := Message{Interaction: "check", Error: err.Error()}
		json.NewEncoder(conn).Encode(answer)
		return
	}

	bytesLocalFile, _ := ioutil.ReadAll(file)
	md5Local := md5.Sum(bytesLocalFile)

	fmt.Println(md5Local)

	answer := Message{
		Interaction: "check",
		Data:        md5Local[:], // arr is an array; arr[:] is the slice of all elements -- STACK OVERFLOW says it best
	}

	json.NewEncoder(conn).Encode(answer)
}
