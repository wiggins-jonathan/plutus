package ingest

import (
    "io/ioutil"
    "path/filepath"
    "fmt"
    "encoding/json"
    "os"

    "gopkg.in/yaml.v3"
)

type Data struct {
    Total   float64
    Tickers map[string]interface{}
}

func Parse(file string) *Data {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("error reading", file, err)
        os.Exit(1)
    }

    var d Data
    extension := filepath.Ext(file)
    switch extension {
    case ".yml", ".yaml"    : err = yaml.Unmarshal(data, &d)
    default                 : err = json.Unmarshal(data, &d)
    }

    if err != nil {
        fmt.Println("Error unmarshalling", file, err)
        os.Exit(1)
    }

    return &d
}
