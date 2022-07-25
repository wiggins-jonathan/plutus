// Unmarshal a json or yaml structure
package ingest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Parse a json or yaml file
func FileParse(file string) map[string]interface{} {
	fileData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading", file, err)
		os.Exit(1)
	}

	// First load the data to map[string]interface{}.
	// This is done because we don't know how many Tickers there will be.
	var data map[string]interface{}
	extension := filepath.Ext(file)
	switch extension {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(fileData, &data)
	default:
		err = json.Unmarshal(fileData, &data)
	}

	if err != nil {
		fmt.Println("Error unmarshalling", file, err)
		os.Exit(1)
	}

	return data
}
