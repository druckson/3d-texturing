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

    //paths := [...]string{
    //    "a/a2/z04.pnm",
    //    "a/a2/z05.pnm",
    //    "a/a2/z06.pnm",
    //    "a/a2/z07.pnm",
    //    "a/a2/z08.pnm",
    //    "a/a2/z09.pnm",
    //    "a/a2/z10.pnm",
    //    "a/a2/z11.pnm",
    //    "a/a2/z12.pnm",
    //    "a/a2/z13.pnm",
    //    "a/a2/z14.pnm",
    //    "a/a2/z15.pnm",
    //    "a/a2/z16.pnm",
    //    "a/a2/z17.pnm",
    //    "a/a2/z18.pnm",
    //    "a/a2/z19.pnm",
    //    "a/a2/z20.pnm",
    //    "a/a2/z21.pnm",
    //    "a/a2/z22.pnm",
    //    "a/a2/z23.pnm",
    //    "a/a2/z24.pnm",
    //    "a/a2/z25.pnm",
    //    "a/b2/z00.pnm",
    //    "a/b2/z01.pnm",
    //    "a/b2/z02.pnm",
    //    "a/b2/z03.pnm",
    //    "a/b2/z04.pnm",
    //    "a/b2/z05.pnm",
    //    "a/b2/z06.pnm",
    //    "a/b2/z07.pnm",
    //    "a/b2/z08.pnm",
    //    "a/b2/z09.pnm",
    //    "a/b2/z10.pnm",
    //    "a/b2/z11.pnm",
    //    "a/b2/z12.pnm",
    //    "a/b2/z13.pnm",
    //    "a/b2/z14.pnm",
    //    "a/b2/z15.pnm",
    //    "a/b2/z16.pnm",
    //    "a/b2/z17.pnm",
    //    "a/b2/z18.pnm",
    //    "a/b2/z19.pnm",
    //    "a/b2/z20.pnm",
    //    "a/b2/z21.pnm",
    //    "a/b2/z22.pnm",
    //    "a/b2/z23.pnm",
    //    "a/b2/z24.pnm",
    //    "a/b2/z25.pnm",
    //    "a/c2/z00.pnm",
    //    "a/c2/z01.pnm",
    //    "a/c2/z02.pnm",
    //    "a/c2/z03.pnm",
    //    "a/c2/z04.pnm",
    //    "a/c2/z05.pnm",
    //    "a/c2/z06.pnm",
    //    "a/c2/z07.pnm",
    //    "a/c2/z08.pnm",
    //    "a/c2/z09.pnm",
    //    "a/c2/z10.pnm",
    //    "a/c2/z11.pnm",
    //    "a/c2/z12.pnm",
    //    "a/c2/z13.pnm",
    //    "a/c2/z14.pnm",
    //    "a/c2/z15.pnm",
    //    "a/c2/z16.pnm",
    //    "a/c2/z17.pnm",
    //    "a/c2/z18.pnm",
    //    "a/c2/z19.pnm",
    //    "a/c2/z20.pnm",
    //    "a/c2/z21.pnm",
    //    "a/c2/z22.pnm",
    //    "a/c2/z23.pnm",
    //    "a/c2/z24.pnm",
    //    "a/c2/z25.pnm",
    //    "a/d2/z00.pnm",
    //    "a/d2/z01.pnm",
    //    "a/d2/z02.pnm",
    //    "a/d2/z03.pnm",
    //    "a/d2/z04.pnm",
    //    "a/d2/z05.pnm",
    //    "a/d2/z06.pnm",
    //    "a/d2/z07.pnm",
    //    "a/d2/z08.pnm",
    //    "a/d2/z09.pnm",
    //    "a/d2/z10.pnm",
    //    "a/d2/z11.pnm",
    //    "a/d2/z12.pnm",
    //    "a/d2/z13.pnm"}

    //sf, _ := readPNMFiles(paths[:])
    //vr := CreateVolumetricRenderer(sf, 0.4, 0.8, 800,
    //                               30.0, 30.0, 15.0,
    //                               90.0, 55.0, 180.0)

    paths := [...]string{
        "a/d2/z15.pnm",
        "a/d2/z16.pnm",
        "a/d2/z17.pnm",
        "a/d2/z18.pnm",
        "a/d2/z19.pnm",
        "a/d2/z20.pnm",
        "a/d2/z21.pnm",
        "a/d2/z22.pnm",
        "a/d2/z23.pnm",
        "a/d2/z24.pnm",
        "a/d2/z25.pnm",
        "a/e2/z00.pnm",
        "a/e2/z01.pnm",
        "a/e2/z02.pnm",
        "a/e2/z03.pnm",
        "a/e2/z04.pnm",
        "a/e2/z05.pnm",
        "a/e2/z06.pnm",
        "a/e2/z07.pnm"}
        //"a/e2/z08.pnm",
        //"a/e2/z09.pnm",
        //"a/e2/z10.pnm",
        //"a/e2/z11.pnm",
        //"a/e2/z12.pnm",
        //"a/e2/z13.pnm",
        //"a/e2/z14.pnm",
        //"a/e2/z15.pnm",
        //"a/e2/z16.pnm",
        //"a/e2/z17.pnm",
        //"a/e2/z18.pnm",
        //"a/e2/z19.pnm",
        //"a/e2/z20.pnm",
        //"a/e2/z21.pnm",
        //"a/e2/z22.pnm",
        //"a/e2/z23.pnm",
        //"a/e2/z24.pnm",
        //"a/e2/z25.pnm",
        //"a/f2/z00.pnm",
        //"a/f2/z01.pnm",
        //"a/f2/z02.pnm",
        //"a/f2/z03.pnm",
        //"a/f2/z04.pnm",
        //"a/f2/z05.pnm",
        //"a/f2/z06.pnm",
        //"a/f2/z07.pnm",
        //"a/f2/z08.pnm",
        //"a/f2/z09.pnm",
        //"a/f2/z10.pnm",
        //"a/f2/z11.pnm",
        //"a/f2/z12.pnm",
        //"a/f2/z13.pnm",
        //"a/f2/z14.pnm",
        //"a/f2/z15.pnm",
        //"a/f2/z16.pnm",
        //"a/f2/z17.pnm",
        //"a/f2/z18.pnm",
        //"a/f2/z19.pnm",
        //"a/f2/z20.pnm",
        //"a/f2/z21.pnm",
        //"a/f2/z22.pnm",
        //"a/f2/z23.pnm",
        //"a/f2/z24.pnm",
        //"a/f2/z25.pnm"}
    sf, _ := readPNMFiles(paths[:])
    vr := CreateVolumetricRenderer(sf, 0.0, 1.0, 800,
                                   30.0, 30.0, 5.0,
                                   90.0, 55.0, 180.0)

    //sf, _ := readVTKFile("astro512.vtk")
    //sf, _ = sf.Half()
    //vr := CreateVolumetricRenderer(sf, 10.0, 15.0, 600,
    //                               30.0, 30.0, 30.0,
    //                               90.0, 75.0, 180.0)

    //sf, _ := readVTKFile("proj6B.vtk")
    //vr := CreateVolumetricRenderer(sf, 0.0, 5.0, 800, 13e+7)

    defer vr.Destroy()

    for {
        vr.Draw()
    }
}

