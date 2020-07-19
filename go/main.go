package main

import (
	"encoding/csv"
	"fmt"
	"frosthage.com/mp3-listaren/formats"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var files []formats.Media

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		media := formats.CreateMedia(path, info)

		files = append(files, media)

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("filer.txt")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	defer file.Close()
	defer writer.Flush()

	for _, f := range files {
		record, err := f.GetRecord()

		if err != nil {
			var v, ok = err.(formats.ErrorMediaFile)

			var record, _ = v.GetRecord()

			if ok {
				writer.Write(strings.Split(record, "\t"))
			} else {
				println(err)
			}
		}

		writer.Write(strings.Split(record, "\t"))
	}
}
