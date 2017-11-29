package main

import (
	"os"
	"fmt"
	"sort"
	"math"
	"net"
	"encoding/json"
)

func ls(conn net.Conn) {
	dir, _ := os.Open(configData.STORAGE_PATH)
	files, _ := dir.Readdir(100)

	sort.Sort(ByName(files))

	var ls string

	for _, file := range files {
		if file.IsDir() {
			subDir, _ := os.Open(file.Name())
			names, _ := subDir.Readdirnames(1000)
			nItems := len(names)

			ls += fmt.Sprintf("\t%v items\t\t%v\n", nItems, file.Name())

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

			ls += fmt.Sprintf("\t%.2f%v\t\t%v\n", size, unit, file.Name())
		}
	}

	json.NewEncoder(conn).Encode(Message{
		Data: []byte(ls),
	})
}