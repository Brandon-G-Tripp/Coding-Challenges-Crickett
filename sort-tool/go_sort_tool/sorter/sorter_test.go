package sorter

import (
    "os"
    "reflect"
    "testing"
)

func TestSortFile(t *testing.T) {
    // Create a temporary file with unsorted content
    content := []byte("banana\napple\ncherry\nbanana\napple\n")
    tmpFile, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpFile.Name())

    if _, err := tmpFile.Write(content); err != nil {
        t.Fatal(err)
    }
    if err := tmpFile.Close(); err != nil {
        t.Fatal(err)
    }

    t.Run("Without -u flag", func(t *testing.T) { 
        sortedLines, err := SortFile(tmpFile.Name(), false)
        if err != nil {
            t.Fatal(err)
        }

        expected := []string{"APPLE", "APPLE", "BANANA", "BANANA", "CHERRY"}
        if !reflect.DeepEqual(sortedLines, expected) {
            t.Errorf("SortFile() = %v, want %v", sortedLines, expected)
        }
    })

    t.Run("With -u flag", func(t *testing.T) {
        sortedLines, err := SortFile(tmpFile.Name(), true)
        if err != nil {
            t.Fatal(err)
        }

        expected := []string{"APPLE", "BANANA", "CHERRY"}
        if !reflect.DeepEqual(sortedLines, expected) {
            t.Errorf("SortFile() = %v, want %v", sortedLines, expected)
        }
    })

}

func TestSortEmptyFile(t *testing.T) {
    // Create an empty temporary file
    tmpFile, err := os.CreateTemp("", "empty")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpFile.Name())

    // Test the SortFile function with an empty file
    sortedLines, err := SortFile(tmpFile.Name(), false)
    if err != nil {
        t.Fatal(err)
    }

    if len(sortedLines) != 0 {
        t.Errorf("SortFile() on empty file should return empty slice, got %v", sortedLines)
    }
}

