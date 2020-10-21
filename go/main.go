package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"frosthage.com/mp3-listaren/formats"
	"golang.org/x/text/encoding/charmap"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
)

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close the paths channel after Walk returns.
		defer close(paths)
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return nil
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

type result struct {
	mediaRecord formats.MediaRecord
	path        string
	err         error
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {

		fileInfo, err := os.Stat(path)
		var record formats.MediaRecord

		if err == nil {
			media := formats.CreateMedia(path, fileInfo)
			record, err = media.GetRecord()
		}

		select {
		case c <- result{record, path, err}:
		case <-done:
			return
		}
	}
}

var root = "."

func main() {

	done := make(<-chan struct{})
	paths, _ := walkFiles(done, root)

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	var wg sync.WaitGroup
	numDigesters := runtime.NumCPU()
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	result := make([]result, 0)

	for r := range c {
		result = append(result, r)
	}

	sort.Slice(result, func(i, j int) bool {
		return strings.Compare(result[i].path, result[j].path) < 0
	})

	file, err := os.Create("filer.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(charmap.Windows1251.NewEncoder().Writer(file))

	defer file.Close()
	defer writer.Flush()

	for _, r := range result {
		if r.err == nil {
			 if err:=writer.Write(r.mediaRecord);err!=nil {
				 println(err)
			 }
		} else {
			var v, ok = r.err.(formats.ErrorMediaFile)
			if ok {
				var record, _ = v.GetRecord()
				if err:= writer.Write(record);err != nil {
					println(err)
				}
			} else {
				println(err)
			}
		}
	}
}

func main1() {

	var files []formats.Media

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

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

	sort.Slice(files, func(i, j int) bool {
		return strings.Compare(files[i].GetPath(), files[j].GetPath()) < 0
	})

	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("filer.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Comma = '\t'

	defer file.Close()
	defer writer.Flush()

	for _, f := range files {
		record, err := f.GetRecord()

		if err != nil {
			var v, ok = err.(formats.ErrorMediaFile)

			var record, _ = v.GetRecord()

			if ok {
				writer.Write(record)
				continue
			} else {
				println(err)
			}
		}

		writer.Write(record)
	}
}
