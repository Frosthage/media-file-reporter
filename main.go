package main

import "C"
import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"frosthage.com/mp3-listaren/formats"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
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

var currentFile = 0

func digester(done <-chan struct{}, paths <-chan string, c chan<- result, fileCount int) {
	for path := range paths {

		fileInfo, err := os.Stat(path)
		var record formats.MediaRecord

		if err == nil {
			media := formats.CreateMedia(path, fileInfo)
			record, err = media.GetRecord()
		}
		select {
		case c <- result{record, path, err}:
			currentFile++
			fmt.Printf("\rFil %d av %d", currentFile, fileCount)
		case <-done:
			return
		}
	}
}

func fileCount(path string) (int, error) {
	i := 0

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			i++
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return i, nil
}

var root = "."

func main() {

	outputFilename := "filer.csv"
	_, err := os.Stat(outputFilename)
	reader := bufio.NewReader(os.Stdin)

	if err == nil {
		exists := true
		for exists {
			fmt.Println("Det finns redan en fil som heter filer.csv i mappen, döp om eller flytta den.")

			reader.ReadRune()
			_, err := os.Stat(outputFilename)
			exists = !os.IsNotExist(err)
		}
	}

	start := time.Now()
	fmt.Println("Fillistaren har startat!")

	fileCount, err := fileCount(root)

	if err != nil {
		fmt.Println(err)
		return
	}

	done := make(<-chan struct{})
	paths, _ := walkFiles(done, root)

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result)
	var wg sync.WaitGroup
	numDigesters := runtime.NumCPU()

	fmt.Printf("Behandlar %d filer samtidigt\n", numDigesters)

	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c, fileCount)
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
	writer := csv.NewWriter(file)

	defer file.Close()
	defer writer.Flush()

	for _, r := range result {
		if r.err == nil {
			if err := writer.Write(r.mediaRecord); err != nil {
				println(err)
			}
		} else {
			var v, ok = r.err.(formats.ErrorMediaFile)
			if ok {
				var record, _ = v.GetRecord()
				if err := writer.Write(record); err != nil {
					println(err)
				}
			} else {
				println(err)
			}
		}
	}

	fmt.Println()
	fmt.Printf("Körningen tog %v", time.Since(start))
	reader.ReadRune()
}
