package main

import ("image"
        "image/png"
        "os")

func readImage(path string) image.Image {
    file, err := os.Open(path)
    if err != nil {
        panic("Failed to read PNM file "+err.Error())
    }

    img, err := png.Decode(file)
    if err != nil {
        panic("Failed to decode PNM file "+err.Error())
    }

    return img
}

func readPNMFiles(paths []string) (*ScalarField, error) {
    img := readImage(paths[0])
    bounds := img.Bounds()
    sf, _ := NewScalarField(bounds.Dx(), bounds.Dy(), len(paths))

    max := float32(0.0)
    for i, path := range(paths) {
        img := readImage(path)
        bounds := img.Bounds()

        for x:=bounds.Min.X; x<bounds.Max.X; x++ {
            for y:=bounds.Min.Y; y<bounds.Max.Y; y++ {
                color := img.At(x, y)
                r, _, _, _ := color.RGBA()
                value := float32(r)/float32(0xFFFF)
                if value > max {
                    max = value
                }
                sf.SetValue(x, y, i, value)
            }
        }
    }
    println("Max ")
    println(max)
    println(sf.width)
    println(sf.height)
    println(sf.depth)
    return sf, nil
}
