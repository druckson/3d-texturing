package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) > 1 {
        conf := CreateConfiguration(os.Args[1])
        sf, _ := readPNMFiles(conf.Files)
        manager := CreateGlfwManager()
        vr := CreateVolumetricRenderer(manager, sf, conf)
        defer vr.Destroy()

        for {
            //vr.Update(0.01)
            //vr.Draw()
            manager.Draw()
        }
    } else {
        fmt.Printf("Please enter the desired project json file\n")
    }
}

