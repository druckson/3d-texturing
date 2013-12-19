package main

func main() {
    vr := CreateVolumetricRenderer()
    defer vr.Destroy()

    for {
        vr.Draw()
    }
}

