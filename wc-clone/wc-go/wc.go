package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func defaultCount(filePath string) (int, int, int, error) {
    var lineCount, wordCount, byteCount int
    var err error

    if filePath == "-" {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            line := scanner.Text()
            lineCount++
            wordCount += len(strings.Fields(line))
            byteCount += len(scanner.Bytes())
        }
        if err := scanner.Err(); err != nil {
            return 0, 0, 0, err
        }
    } else {
        lineCount, err = countLines(filePath)
        if err != nil {
            return 0, 0, 0, err
        } 

        wordCount, err = countWords(filePath)
        if err != nil {
            return 0, 0, 0, err
        } 

        byteCount, err = countBytes(filePath)
        if err != nil {
            return 0, 0, 0, err
        } 
    }

    return lineCount, wordCount, byteCount, nil
}

func countChars(filePath string) (int, error) {
    var data []byte
    var err error
    if filePath == "-" {
        data, err = io.ReadAll(os.Stdin)
    } else {
        data, err = os.ReadFile(filePath)
    } 
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
    var scanner *bufio.Scanner
    if filePath == "-" {
        scanner = bufio.NewScanner(os.Stdin)
    } else {
        file, err := os.Open(filePath)
        if err != nil {
            return 0, err
        } 
        defer file.Close()
        scanner = bufio.NewScanner(file)
    }

    scanner.Split(bufio.ScanWords)
    count := 0 
    for scanner.Scan() {
        count++
    } 
    return count, nil
} 

func countBytes(filePath string) (int, error) {
    var data []byte
    var err error
    if filePath == "-" {
        data, err = io.ReadAll(os.Stdin)
    } else {
        data, err = os.ReadFile(filePath)
    }
    if err != nil {
        return 0, err
    }
    return len(data), nil
} 

func countLines(filePath string) (int, error) {
    var scanner *bufio.Scanner
    if filePath == "-" {
        scanner = bufio.NewScanner(os.Stdin)
    } else {
        file, err := os.Open(filePath)
        if err != nil {
            return 0, err
        }
        defer file.Close()
        scanner = bufio.NewScanner(file)
    }
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

    var filePath string
    if flag.NArg() == 0 {
        filePath = "-"
    } else {
        filePath = flag.Arg(0)
    } 

    if !*countBytesFlag && !*countLinesFlag && !*countWordsFlag && !*countCharsFlag {
        lineCount, wordCount, byteCount, err := defaultCount(filePath)
        if err != nil {
            fmt.Println("Error: Could not read input")
            os.Exit(1)
        } 
        fmt.Printf("%d %d %d\n", lineCount, wordCount, byteCount)
        return
    }

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
