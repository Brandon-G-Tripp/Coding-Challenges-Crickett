package sorter

import (
	"bufio"
	"os"
	"strings"

    "github.com/Brandon-G-Tripp/go_sort_tool/sorter/algorithms"
)

func SortFile(filename string, unique bool, algorithm string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(strings.ToUpper(scanner.Text()))
        if line != "" {
            lines = append(lines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    switch algorithm {
    case "merge":
        algorithms.MergeSort(lines)
    case "quick":
        algorithms.QuickSort(lines)
    case "heap":
        algorithms.HeapSort(lines)
    case "radix":
        algorithms.RadixSort(lines)
    default:
        // for now, use MergeSort as default
        algorithms.QuickSort(lines)
    }

    if unique {
        return removeDuplicates(lines), nil
    }

    return lines, nil
}

func removeDuplicates(lines []string) []string {
    var uniqueLines []string
    for i, line := range lines {
        if i == 0 || line != lines[i-1] {
            uniqueLines = append(uniqueLines, line)
        }
    }
    return uniqueLines
}
