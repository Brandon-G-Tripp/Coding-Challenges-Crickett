package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func countBytes(filePath string) (int, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return 0, err
    }
    return len(data), nil
} 

func main() {
    countBytesFlag := flag.Bool("c", false, "Count bytes")
    flag.Parse()

    if flag.NArg() == 0 {
        fmt.Println("Please provide a file")
        os.Exit(1)
    } 

    filePath := flag.Arg(0)

    if *countBytesFlag {
        count, err := countBytes(filePath)
        if err != nil {
            fmt.Printf("Error: Could not open file '%s'\n", filePath)
            os.Exit(1)
        }
        fmt.Printf("%d %s\n", count, filePath)
    } else {
        fmt.Println("Error: Missing -c flag")
        os.Exit(1)
    }
} 
