package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func countChars(filePath string) (int, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return 0, err
    } 

    count := 0
    for range string(data) {
        count++
    } 

    return count, nil
}

func countWords(filePath string) (int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 0, err
    } 
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    count := 0 
    for scanner.Scan() {
        count++
    } 
    return count, nil
} 

func countBytes(filePath string) (int, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return 0, err
    }
    return len(data), nil
} 

func countLines(filePath string) (int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }
    return count, nil
} 

func main() {
    countBytesFlag := flag.Bool("c", false, "Count bytes")
    countLinesFlag := flag.Bool("l", false, "Count lines")
    countWordsFlag := flag.Bool("w", false, "Count words")
    countCharsFlag := flag.Bool("m", false, "Count characters")
    flag.Parse()

    if flag.NArg() == 0 {
        fmt.Println("Please provide a file")
        os.Exit(1)
    } 

    filePath := flag.Arg(0)

    if *countCharsFlag {
        count, err := countChars(filePath)
        if err != nil {
            fmt.Printf("Error: Could not open file '%s'\n", filePath)
            os.Exit(1)
        } 
        fmt.Printf("%d %s\n", count, filePath)
    } else if *countBytesFlag {
        count, err := countBytes(filePath)
        if err != nil {
            fmt.Printf("Error: Could not open file '%s'\n", filePath)
            os.Exit(1)
        }
        fmt.Printf("%d %s\n", count, filePath)
    } else if *countLinesFlag {
        count, err := countLines(filePath)
        if err != nil {
            fmt.Printf("Error: Could not open file '%s'\n", filePath)
            os.Exit(1)
        } 
        fmt.Printf("%d %s\n", count, filePath)
    } else if *countWordsFlag {
        count, err := countWords(filePath)
        if err != nil {
            fmt.Printf("Error: Could not open file '%s'\n", filePath)
            os.Exit(1)
        } 
        fmt.Printf("%d %s\n", count, filePath)
    } else {
        fmt.Println("Error: Missing -c, -l, -m,  or -w flag")
        os.Exit(1)
    }
} 
