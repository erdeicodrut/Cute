package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func config(_ *cli.Context) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("IP: ")
	initData.IP, _ = reader.ReadString('\n')

	fmt.Print("PORT: ")
	initData.PORT, _ = reader.ReadString('\n')

	initData.IP = strings.Trim(initData.IP, "\n")
	initData.IP = strings.Trim(initData.IP, " ")
	initData.PORT = strings.Trim(initData.PORT, "\n")
	initData.PORT = strings.Trim(initData.PORT, " ")

	file, err := os.Open("init.json")
	if err != nil {
		file, err = os.Create("init.json")
		fmt.Printf("Err: %v", err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(meta{initData.IP, initData.PORT})

}
