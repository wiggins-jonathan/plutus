package ingest

import (
    "io/ioutil"
    "path/filepath"
     "fmt"

    "gopkg.in/yaml.v3"
)

type Data struct {
    Total   string `yaml:"total"`
    Tickers map[string]struct {
        CurrentAmount   float64 `yaml:"current"`
        DesiredPercent  float64 `yaml:"desired"`
    }
}

func Parse(file string) *Data {
    var d Data
    extension := filepath.Ext(file)
    switch extension {
    case ".yml", ".yaml"    : d = yml(file)
    default                 : d = yml(file)
    }

    return &d
}

func yml(file string) Data {
    var d Data
    data, err := ioutil.ReadFile(file)
    if err != nil {fmt.Println("error reading ", file)}

    if err := yaml.Unmarshal(data, &d); err != nil {
        fmt.Println("Error unmarshalling ", file)
    }

    return d
}
