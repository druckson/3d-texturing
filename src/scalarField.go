package main

import "errors"
import "math"
import "fmt"

type ScalarField struct {
    width int
    height int
    depth int
    data []float32
}

func NewScalarField(w int, h int, d int) (*ScalarField, error) {
    if w < 0 || h < 0 || d < 0 {
        return nil, errors.New("ScalarField: Invalid dimensions")
    }

    sf := new(ScalarField)
    sf.width = int(math.Pow(2.0, math.Floor(math.Log2(float64(w)))))
    sf.height = int(math.Pow(2.0, math.Floor(math.Log2(float64(h)))))
    sf.depth = int(math.Pow(2.0, math.Floor(math.Log2(float64(d)))))
    sf.data = make([]float32, w*h*d)
    return sf, nil
}

func (sf *ScalarField) SetData(data []float32) error {
    //sf.data = data
    fmt.Printf("Len: %v\n", len(data))
    for i, num := range(data) {
        //fmt.Printf("num: %v\n", num)
        sf.data[i] = num
    }
    return nil
}

func (sf *ScalarField) index(x int, y int, z int) (int, error) {
    if  x >= 0 && x < sf.width &&
        y >= 0 && y < sf.height &&
        z >= 0 && z < sf.depth {
        //return x*sf.height*sf.depth + y*sf.depth + z, nil
        //return z*sf.width*sf.height + y*sf.depth + x, nil
        return z*sf.width*sf.height + y*sf.width + x, nil
    }
    return 0.0, errors.New("ScalarField: Coordinates out of bounds")
}

func (sf *ScalarField) GetValue(x int, y int, z int) float32 {
    index, err := sf.index(x, y, z)

    if err != nil {
        panic(err)
    }

    return sf.data[index]
}

func (sf *ScalarField) SetValue(x int, y int, z int, value float32) error {
    index, err := sf.index(x, y, z)

    if err != nil {
        return err
    }

    sf.data[index] = value
    return nil
}

func (sf *ScalarField) Half() (*ScalarField, error) {
    new_sf, _ := NewScalarField(sf.width/2, sf.height/2, sf.depth/2)

    for x := 0; x<new_sf.width; x++ {
        for y := 0; y<new_sf.height; y++ {
            for z := 0; z<new_sf.depth; z++ {
                new_sf.SetValue(x, y, z, (
                    sf.GetValue(2*x+0, 2*y+0, 2*z+0) +
                    sf.GetValue(2*x+0, 2*y+0, 2*z+1) +
                    sf.GetValue(2*x+0, 2*y+1, 2*z+0) +
                    sf.GetValue(2*x+0, 2*y+1, 2*z+1) +
                    sf.GetValue(2*x+1, 2*y+0, 2*z+0) +
                    sf.GetValue(2*x+1, 2*y+0, 2*z+1) +
                    sf.GetValue(2*x+1, 2*y+1, 2*z+0) +
                    sf.GetValue(2*x+1, 2*y+1, 2*z+1))/8.0)
            }
        }
    }

    return new_sf, nil
}
