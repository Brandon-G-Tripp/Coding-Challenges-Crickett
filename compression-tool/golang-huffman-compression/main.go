package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a file path as an argument")
        return
    }

    filePath := os.Args[1]
    contents, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return 
    }

    input := string(contents)
    frequencies := CountCharacterFrequencies(input)
    printFrequencies(frequencies)
}

func printFrequencies(frequencies map[rune]int) {
    for char, count := range frequencies {
        fmt.Printf("%s: %d\n", string(char), count)
    }
}
