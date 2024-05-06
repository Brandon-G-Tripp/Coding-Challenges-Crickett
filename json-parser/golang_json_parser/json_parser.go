package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run json_parser.go <json_file>")
        os.Exit(1)
    }

    filename := os.Args[1]
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("Error reading file: %s\n", err)
        os.Exit(1)
    } 

    json := string(content)
    if isValidJSON(json) {
        fmt.Println("Valid JSON")
        os.Exit(0)
    } else {
        fmt.Println("Invalid JSON")
        os.Exit(1)
    }
} 

func isValidJSON(json string) bool {
    json = strings.TrimSpace(json)
    return len(json) >= 2 && json[0] == '{' && json[len(json) - 1] == '}'
} 


