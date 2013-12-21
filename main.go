package main

//import "fmt"

func main() {
    //sf, _ := NewScalarField(100, 100, 100)
    //sf, _=
    //val := sf.GetValue(0, 0, 0)
    //fmt.Printf("Hello %v\n", val)

    sf, _ := readVTKFile("astro64.vtk")
    vr := CreateVolumetricRenderer(sf, 10.0, 15.0, 9e+7)

    //sf, _ := readVTKFile("proj6B.vtk")
    //vr := CreateVolumetricRenderer(sf, 0.0, 5.0, 13e+7)

    defer vr.Destroy()

    for {
        vr.Draw()
    }
}

