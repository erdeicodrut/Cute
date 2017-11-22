package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"sort"
)

type ByName []os.FileInfo
func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }


func main() {
	app := cli.NewApp()

	app.Name = "Cute"
	app.Usage = "A simple cloud storage kind of stuff"
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "lang, l",
			Usage: "language for the greeting",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("lang") == false {
			fmt.Println("Nu este flag")
		} else {
			fmt.Println("da")
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "Pushes a file",
			Action: func(c *cli.Context) error {
				fmt.Println("Pushed the file")
				return nil
			},
		},

		{
			Name:  "pull",
			Usage: "Pulls a file",
			Action: func(c *cli.Context) error {
				fmt.Println("Pulled the file")
				return nil
			},
		},

		{
			Name:  "ls",
			Usage: "Lists all the files",
			Action: func(c *cli.Context) error {
				ls, _ := os.Open("./")
				files, _ := ls.Readdir(100)

				sort.Sort(ByName(files))

				for _, file := range files {
					fmt.Printf("Name: %v\tSize: %v\n", file.Name(), file.Size())
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
