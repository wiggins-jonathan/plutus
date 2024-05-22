// Unmarshal a json or yaml structure
package ingest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// Parse a json or yaml file
func FileParse(file string) (map[string]any, error) {
	fileData, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("error reading %s: %w", file, err)
	}

	// First load the data to map[string]any.
	// This is done because we don't know how many Tickers there will be.
	var data map[string]any
	extension := filepath.Ext(file)
	switch extension {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &data)
	default:
		err = json.Unmarshal(fileData, &data)
	}

	if err != nil {
		fmt.Errorf("Error unmarshalling %s: %w", file, err)
	}

	return data, nil
}

// ReadFiles walks a dir & reads files with extension concurrently into memory
func ReadFiles(path, extension string) {
	numCPU := runtime.NumCPU()
	numRoutines := numCPU * 2

	files := make(chan string, 100)
	results := make(chan []byte, 100)

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	// Start goroutines to process files
	for i := 0; i < numRoutines; i++ {
		go func(files <-chan string, results chan<- []byte) {
			defer wg.Done()
			for file := range files {
				data, err := ioutil.ReadFile(file)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", file, err)
					continue
				}
				results <- data
			}
		}(files, results)
	}

	// Start a goroutine to process results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Walk through directory and send files to goroutines
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") { // skip hidden dirs & files
			return filepath.SkipDir
		}

		if !info.IsDir() && filepath.Ext(path) == extension {
			files <- path
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}

	close(files)

	// Process data from results
	for data := range results {
		fmt.Println(string(data))
	}
}
