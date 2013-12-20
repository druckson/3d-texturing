package main

import ("strconv"
        "os"
        "strings"
        "encoding/binary"
        "bytes")

func readFile(path string) []byte {
    file, err := os.Open(path)
    if err != nil {
        panic("Failed to open the path")
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic("Failed to find file size")
    }

    data := make([]byte, stat.Size())
    _, err = file.Read(data)
    if err != nil {
        panic("Failed to read file")
    }

    return data
}

func readVTKFile(path string) ([]float32, []float32) {
    data := strings.Split(strings.Join(strings.Split(string(readFile(path)), "\n"), ""), " ");

    nums := make([]float32, len(data))

    for i := 0; i<len(data); i++ {
        num, _ := strconv.ParseFloat(data[i], 32)
        nums[i] = float32(num)
    }

    return nil, nums
}

func readVTKBinaryFile(path string) ([]float32, []float32) {
    data := bytes.NewBuffer(readFile(path))
    nums := make([]float32, data.Len()/4)

    for i := 0; i<len(nums); i++ {
        err := binary.Read(data, binary.BigEndian, &nums[i])
        if err != nil {
            panic("Failed to read file")
        }
    }

    return nil, nums
}
