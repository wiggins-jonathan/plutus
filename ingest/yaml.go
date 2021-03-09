package ingest

import (
    "io/ioutil"
    "path/filepath"
     "fmt"
     "encoding/json"

    "gopkg.in/yaml.v3"
)

type Data struct {
    Total   float64 `yaml:"total"`
    Tickers map[string]struct {
        CurrentAmount   float64 `yaml:"current"`
        DesiredPercent  float64 `yaml:"desired"`
    }
}

func Parse(file string) *Data {
    data, err := ioutil.ReadFile(file)
    if err != nil { fmt.Println("error reading", file, err) }

    var d Data
    extension := filepath.Ext(file)
    switch extension {
    case ".yml", ".yaml"    : err = yaml.Unmarshal(data, &d)
    default                 : err = json.Unmarshal(data, &d)
    }

    if err != nil { fmt.Println("Error unmarshalling", file, err) }

    return &d
}
