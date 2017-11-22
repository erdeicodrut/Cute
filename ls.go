package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
)

func ls(c *cli.Context) error {
	ls, _ := os.Open("./")
	files, _ := ls.Readdir(100)

	sort.Sort(ByName(files))

	for _, file := range files {
		fmt.Printf("\tSize: %.2f kb\t %v\n", float32(file.Size())/1024, file.Name())
		// TODO format the the size in b/kb/mb/gb and take care of tabs

	}
	return nil
}
