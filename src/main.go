package main

import (
    "os"
    "fmt"
)

func main() {
    //vr := CreateRenderer("../data/mom1.json")
    fmt.Printf("Args: %s\n", os.Args[1])
    vr := CreateRenderer(os.Args[1])
    defer vr.Destroy()

    for {
        vr.Update(0.01)
        vr.Draw()
    }
}

