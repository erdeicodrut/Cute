package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
	"math"
)

func ls(c *cli.Context) error {
	ls, _ := os.Open("./")
	files, _ := ls.Readdir(100)

	sort.Sort(ByName(files))

	for _, file := range files {
		if file.IsDir() {
			subDir, _ := os.Open(file.Name())
			names, _ := subDir.Readdirnames(1000)
			nItems := len(names)
			
			fmt.Printf("\t%v items\t\t\t\t%v\n", nItems, file.Name())

		} else {
			var size float64
			var unit string

			if file.Size() >= int64(math.Pow(1000, 3)) {
				size = float64(file.Size()) / math.Pow(1000, 3)
				unit = "GB"
			} else if file.Size() >= int64(math.Pow(1000, 2)) {
				size = float64(file.Size()) / math.Pow(1000, 2)
				unit = "MB"
			} else if file.Size() >= 1000 {
				size = float64(file.Size()) / 1000
				unit = "KB"
			} else {
				size = float64(file.Size())
				unit = "B"
			}

			fmt.Printf("\tSize: %.2f%v\t%v\n", size, unit, file.Name())
		}
	}
	return nil
}
