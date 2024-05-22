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
