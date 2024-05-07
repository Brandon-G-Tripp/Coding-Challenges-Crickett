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

    if !strings.HasPrefix(json, "{") || !strings.HasSuffix(json, "}") {
        return false
    }

    if json == "{}" {
        return true
    }

    content := json[1 : len(json) - 1]
    pairs := strings.Split(content, ",")

    for _, pair := range pairs {
        parts := strings.SplitN(pair, ":", 2)
        if len(parts) != 2 {
            return false
        }

        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])

        if !isValidString(key) || !isValidString(value) {
            return false
        }
    }

    return true
} 

func isValidString(str string) bool {
    return len(str) >= 2 && strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"")
}
