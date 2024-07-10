package main

import (
	"flag"
	"fmt"
	"go_sort_tool/sorter"
	"os"
)

var exit = os.Exit

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [-u] <filename>\n", os.Args[0])
    }
    uniqueFlag := flag.Bool("u", false, "Remove duplicate lines")
    flag.Parse()

    args := flag.Args()
    if len(args) != 1 {
        flag.Usage()
        exit(1)
        return
    }

    filename := args[0]
    sortedLines, err := sorter.SortFile(filename, *uniqueFlag)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        exit(1)
        return 
    }

    for _, line := range sortedLines {
        fmt.Println(line)
    }
}
