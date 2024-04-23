package main

import (
    "os"
    "testing"
)

func TestCountLines(t *testing.T) {
    file, err := os.Create("test.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary file: %v", err)
    } 
    defer file.Close()
    defer os.Remove("test.txt")

    content := "Line 1\nLine 2\nLine 3\n"
    _, err = file.WriteString(content)
    if err != nil {
        t.Fatalf("Failed to write to temporary file: %v", err)
    } 

    count, err := countLines("test.txt")
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    } 

    expectedCount := 3
    if count != expectedCount {
        t.Errorf("Expected line count %d, but got %d", expectedCount, count)
    } 

}

func TestCountLinesFileNotFound(t *testing.T) {
    _, err := countLines("nonexistent.txt")

    if err == nil {
        t.Error("Expected an error, but got nil")
    } 
} 

func TestCountBytes(t *testing.T) {
    // Create a temp file with sample input
    file, err := os.Create("test.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary file: %v", err)
    }
    defer file.Close()
    defer os.Remove("test.txt")

    content := []byte("Sample content")
    _, err = file.Write(content)
    if err != nil {
        t.Fatalf("Failed to write to temproary file: %v", err)
    } 

    // Call the countBytes function
    count, err := countBytes("test.txt")
    if err != nil {
        t.Fatalf("Unexpted error: %v", err)
    }

    // Assert the expected byte count
    if count != len(content) {
        t.Errorf("Expected byte count %d, but got %d", len(content), count)
    } 
} 

func TestCountBytesFileNotFound(t *testing.T) {
    // Call the countBytes function with a non-existent file
    _, err := countBytes("nonexistent.txt")

    // Assert that the expected error is returned
    if err == nil {
        t.Error("Expected an error, but got nil")
    } 
} 
