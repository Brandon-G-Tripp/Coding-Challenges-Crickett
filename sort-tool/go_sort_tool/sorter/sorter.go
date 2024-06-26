package sorter

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func SortFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, strings.ToUpper(strings.TrimSpace(scanner.Text())))
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    sort.Strings(lines)

    // Remove duplicates
    var uniqueLines []string
    for i, line := range lines {
        if i == 0 || line != lines[i-1] {
            uniqueLines = append(uniqueLines, line)
        }
    }

    return uniqueLines, nil
}
