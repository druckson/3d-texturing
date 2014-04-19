package main

import "os"

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
