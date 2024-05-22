package main

import (
    "fmt"
    "os"

    "github.com/Brandon-G-Tripp/golang-cut-tool/cut"
)

func main() {
    if len(os.Args) != 3 || os.Args[1] != "-f2" {
        fmt.Println("Usage: cut -f2 <file>")
        os.Exit(1)
    }

    filePath := os.Args[2]
    if err := cut.CutSecondField(filePath); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
