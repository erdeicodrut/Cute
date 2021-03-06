package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
	"net"
	"os"
	"strings"
)

type Config struct {
	PORT         string `json:"port"`
	STORAGE_PATH string `json:"storage_path"`
}

type Message struct {
	Interaction string `json:"interaction"`
	Name        string `json:"name"`
	Data        []byte `json:"data"`
	Error       string `json:"error"`
	Date        uint64 `json:"date"`
}

var configData = Config{"", ""}

var dbC chan *sql.DB

func main() {
	dbC = make(chan *sql.DB, 1)
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	dbC <- db

	app := cli.NewApp()

	configFile, err := os.Open("config.json")
	if err != nil {
		config(nil)
	}

	json.NewDecoder(configFile).Decode(&configData)
	configFile.Close()

	app.Name = "Cute (server)"
	app.Usage = "A simple cloud storage kind of stuff"
	app.HideVersion = true

	app.Commands = []cli.Command{
		{
			Name:   "config",
			Usage:  "Configure Cute server",
			Action: config,
		},
	}

	app.Action = func(c *cli.Context) error {
		ln, err := net.Listen("tcp", ":"+configData.PORT)
		if err != nil {
			fmt.Println("Invalid port, please reconfigure.")
			return nil
		}
		defer ln.Close()

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Connection refused.")
				continue
			}
			go handleConnection(conn)
		}
	}

	app.Run(os.Args)
}

func handleConnection(conn net.Conn) {
	message := Message{}
	json.NewDecoder(conn).Decode(&message)

	switch message.Interaction {
	case "push":
		receive(message)
	case "pull":
		send(message, conn)
	case "ls":
		ls(conn)
	case "check":
		check(message.Name, conn)
	case "rm":
		rm(message)
	}
}

type ByName []os.FileInfo

func config(c *cli.Context) {
	file, err := os.Create("config.json")
	if err != nil {
		fmt.Printf("Couldn't create 'config.json' file because %v\n", err)
	}

	defer file.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("PORT: ")
	PORT, _ := reader.ReadString('\n')

	fmt.Print("STORAGE_PATH: ")
	STORAGE_PATH, _ := reader.ReadString('\n')

	PORT = strings.Trim(PORT, "\n"+" ")
	STORAGE_PATH = strings.Trim(STORAGE_PATH, "\n"+" ")

	// Append '/' at the end of path if there isn't one
	if STORAGE_PATH[len(STORAGE_PATH)-1] != '/' {
		STORAGE_PATH += "/"
	}

	// Create the directories
	os.MkdirAll(STORAGE_PATH, 0755)

	configData = Config{PORT, STORAGE_PATH}

	json.NewEncoder(file).Encode(configData)

	fmt.Println(file)
}

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }
