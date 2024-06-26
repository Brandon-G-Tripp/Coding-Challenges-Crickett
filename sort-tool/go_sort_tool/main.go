package main

import (
    "fmt"
    "os"
    "go_sort_tool/sorter"
)

var exit = os.Exit

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <filename>\n", os.Args[0])
        exit(1)
        return 
    }

    filename := os.Args[1]
    sortedLines, err := sorter.SortFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        exit(1)
        return 
    }

    for _, line := range sortedLines {
        fmt.Println(line)
    }
}
