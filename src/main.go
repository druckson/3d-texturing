package main

import (
    "os"
    "fmt"
)

func main() {
    //vr := CreateRenderer("../data/mom1.json")
    if len(os.Args) > 1 {
        //vr := CreateRenderer(os.Args[1])
        conf := CreateConfiguration(os.Args[1])
        sf, _ := readPNMFiles(conf.Files)
        vr := CreateVolumetricRenderer(sf, conf)
        defer vr.Destroy()

        for {
            vr.Update(0.01)
            vr.Draw()
        }
    } else {
        fmt.Printf("Please enter the desired project json file\n")
    }
}

