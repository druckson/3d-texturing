package main

//import "fmt"

func main() {
    //sf, _ := NewScalarField(100, 100, 100)
    //sf, _=
    //val := sf.GetValue(0, 0, 0)
    //fmt.Printf("Hello %v\n", val)

    //sf, _ := readVTKFile("astro64.vtk")
    //vr := CreateVolumetricRenderer(sf, 10.0, 15.0, 300,
    //                               3.02e+7, 3.02e+7, 3.02e+7,
    //                               9e+7, 7.5e+7, 1.8e+8)

    //sf, _ := readVTKFile("astro64.vtk")
    //vr := CreateVolumetricRenderer(sf, 10.0, 15.0, 300,
    //                               30.0, 30.0, 30.0,
    //                               90.0, 75.0, 180.0)

    sf, _ := readVTKFile("astro512.vtk")
    sf, _ = sf.Half()
    sf, _ = sf.Half()
    vr := CreateVolumetricRenderer(sf, 10.0, 15.0, 600,
                                   30.0, 30.0, 30.0,
                                   90.0, 75.0, 180.0)

    //sf, _ := readVTKFile("proj6B.vtk")
    //vr := CreateVolumetricRenderer(sf, 0.0, 5.0, 800, 13e+7)

    defer vr.Destroy()

    for {
        vr.Draw()
    }
}

