package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"frosthage.com/mp3-listaren/formats"
	"os"
	"path/filepath"
	"runtime"
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
				return err
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
	media formats.Media
	err   error
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {

		fileInfo, err := os.Stat(path)
		var media formats.Media

		if err == nil {
			media = formats.CreateMedia(path, fileInfo)
		}

		select {
		case c <- result{media, err}:
		case <-done:
			return
		}
	}
}


func main() {

	done := make(<-chan struct{})
	paths, _ := walkFiles(done, ".")

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

	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	for r := range c {
		if r.err == nil {
			r, _ := r.media.GetRecord()
			writer.Write(strings.Split(r, "\t"))
		}
	}
}


func main2() {

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

