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
        controller := CreateCameraController(manager, vr)
        defer vr.Destroy()

        for {
            controller.Update(0.01)
            manager.Draw()
        }
    } else {
        fmt.Printf("Please enter the desired project json file\n")
    }
}

