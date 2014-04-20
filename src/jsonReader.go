package main

import (
    "os"
    "path"
    "fmt"
    "encoding/json"
)

type Project struct {
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

func CreateRenderer(projectFile string) *VolumetricRenderer {
    fileinfo, _ := os.Stat(projectFile)
    file, _ := os.Open(projectFile)
    size := fileinfo.Size()
    data := make([]byte, size)
    file.Read(data)

    fmt.Printf("Size: %d\n", size)

    var proj Project
    json.Unmarshal(data, &proj)
    fmt.Printf("Test1: %d\n", len(proj.Paths))
    fmt.Printf("Test1: %f\n", proj.Max)

    base := path.Dir(projectFile)
    fmt.Printf("Test1: %s\n", base)
    for i, filePath := range proj.Paths {
        proj.Paths[i] = path.Join(base, filePath)
    }

    sf, _ := readPNMFiles(proj.Paths)
    fmt.Printf("Test2\n")
    vr := CreateVolumetricRenderer(sf,
        proj.Min, proj.Max, proj.Samples,
        proj.Width, proj.Height, proj.Depth,
        proj.Dist, proj.Near, proj.Far)
    fmt.Printf("Test3\n")
    return vr
}
