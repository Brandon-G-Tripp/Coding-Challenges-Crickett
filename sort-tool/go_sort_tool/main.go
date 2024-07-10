package main

import (
	"flag"
	"fmt"
	"github.com/Brandon-G-Tripp/go_sort_tool/sorter"
	"os"
)

func run(args []string) error {
    flags := flag.NewFlagSet(args[0], flag.ExitOnError)
    uniqueFlag := flags.Bool("u", false, "Remove duplicate lines")
    algorithmFlag := flags.String("a", "quick", "Sorting algorithm to use (merge, quick, heap, radix)")

    if err := flags.Parse(args[1:]); err != nil {
        return err
    }

    if flags.NArg() != 1 {
        return fmt.Errorf("Usage: %s [-u] [-a algorithm] <filename>", args[0])
    }

    filename := flags.Arg(0)

    // Validate and set the sorting algorithm
    algorithm := *algorithmFlag
    if !isValidAlgorithm(algorithm) {
        fmt.Fprintf(os.Stderr, "Warning: Unknown sorting algorithm '%s'. Defaulting to quick sort.\n", algorithm)
        algorithm = "quick"
    }

    sortedLines, err := sorter.SortFile(filename, *uniqueFlag, algorithm)
    if err != nil {
        return err
    }

    for _, line := range sortedLines {
        fmt.Println(line)
    }

    return nil
}

func isValidAlgorithm(algo string) bool {
    validAlgorithms := []string{"merge", "quick", "heap", "radix"}
    for _, validAlgo := range validAlgorithms {
        if algo == validAlgo {
            return true
        }
    }
    return false
}

func main() {
    if err := run(os.Args); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
