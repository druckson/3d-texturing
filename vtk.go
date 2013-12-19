package main

import ("strconv"
        "os"
        "strings")

func readFile(path string) string {
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

    return string(data)
}

func readVTKFile(path string) ([]float32, []float32) {
    data := strings.Split(strings.Join(strings.Split(readFile(path), "\n"), " "), " ");

    nums := make([]float32, len(data))

    for i := 0; i<len(data); i++ {
        //fmt.Printf("%v\n", data[i])
        num, _ := strconv.ParseFloat(data[i], 32)
        nums[i] = float32(num)
    }

    return nums, nil
}
