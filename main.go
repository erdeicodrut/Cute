package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("wall10.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	byt, _ := ioutil.ReadAll(file)

	clone, _ := os.Create("Groaznic.jpg")
	clone.Write(byt)

	defer clone.Close()
}
