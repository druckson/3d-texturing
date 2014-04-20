package main

import (
    "os"
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
    Paths []string
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

    base := path.Dir(configFile)
    for i, filePath := range conf.Paths {
        conf.Paths[i] = path.Join(base, filePath)
    }

    return conf
}

func CreateRenderer(configFile string) *VolumetricRenderer {
    conf := CreateConfiguration(configFile);

    sf, _ := readPNMFiles(conf.Paths)
    vr := CreateVolumetricRenderer(sf,
        conf.Min, conf.Max, conf.Samples,
        conf.Width, conf.Height, conf.Depth,
        conf.Dist, conf.Near, conf.Far)

    return vr
}
