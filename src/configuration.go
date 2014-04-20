package main

import (
    "os"
    "fmt"
    "path"
    "encoding/json"
)

type Configuration struct {
    Min float32
    Max float32
    Samples int
    Width float32
    Height float32
    Depth float32
    Dist float32
    Near float32
    Far float32
    Files []string
}

func CreateConfiguration(configFile string) *Configuration {
    fileinfo, _ := os.Stat(configFile)
    file, _ := os.Open(configFile)
    defer file.Close()

    size := fileinfo.Size()
    data := make([]byte, size)
    file.Read(data)

    var conf *Configuration
    json.Unmarshal(data, &conf)

    fmt.Printf("File: %f\n", conf.Min)
    base := path.Dir(configFile)
    for i, filePath := range conf.Files {
        conf.Files[i] = path.Join(base, filePath)
    }

    return conf
}
