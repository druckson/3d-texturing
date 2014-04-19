package main

import ("strconv"
        "strings"
        "encoding/binary"
        "bytes"
        "fmt")

func readVTKTextData(lines []string) ([]float32) {
    data := strings.Split(strings.TrimRight(strings.Join(lines, ""), " "), " ");
    nums := make([]float32, len(data))

    for i := 0; i<len(data); i++ {
        num, _ := strconv.ParseFloat(data[i], 32)
        nums[i] = float32(num)
    }

    return nums
}

func readVTKBinaryData(lines []string) ([]float32) {
    data := bytes.NewBuffer([]byte(strings.Join(lines, "\n")))
    nums := make([]float32, data.Len()/4)

    for i := 0; i<len(nums); i++ {
        err := binary.Read(data, binary.BigEndian, &nums[i])
        if err != nil {
            panic("Failed to read file")
        }
    }

    return nums
}

func readVTKFile(path string) (*ScalarField, error) {
    lines := strings.Split(string(readFile(path)), "\n");

    var i, w, h, d int
    var line string
    var bin bool
    for i, line = range lines {
        if line == "ASCII" {
            fmt.Printf("Ascii file\n")
            bin = false
        } else if line == "BINARY" {
            fmt.Printf("Binary file\n")
            bin = true
        } else if len(line) >= 12 && line[:12] == "LOOKUP_TABLE" {
            fmt.Printf("Lookup Table\n")
            break
        } else if len(line) >= 10 &&  line[:10] == "DIMENSIONS" {
            dims := strings.Split(line, " ")
            fmt.Printf("Dimensions: %v\n", dims)
            x, _ := strconv.ParseInt(dims[1], 10, 32)
            y, _ := strconv.ParseInt(dims[2], 10, 32)
            z, _ := strconv.ParseInt(dims[3], 10, 32)
            w = int(x)
            h = int(y)
            d = int(z)
        }
    }

    var data []float32
    if bin {
        data = readVTKBinaryData(lines[i+1:])
    } else {
        data = readVTKTextData(lines[i+1:])
    }

    sf, _ := NewScalarField(w, h, d)
    sf.SetData(data)

    return sf, nil
}
