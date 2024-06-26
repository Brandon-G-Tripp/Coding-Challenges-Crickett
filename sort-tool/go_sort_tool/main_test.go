package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestCliWithWordsFile(t *testing.T) {
    filename := "../words.txt"

    // check if the file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        t.Fatalf("words.txt file not found in the parent directory. Please ensure it is created before running the test.")
    }

    // run the CLI command
    cmd := exec.Command("go", "run", "main.go", filename)
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        t.Fatalf("Failed to run CLI command: %v", err)
    }

    // Pipe the output through uniq and head
    uniqCmd := exec.Command("uniq")
    uniqCmd.Stdin = strings.NewReader(out.String())
    var uniqOut bytes.Buffer
    uniqCmd.Stdout = &uniqOut

    err = uniqCmd.Run() 
    if err != nil {
        t.Fatalf("Failed to run uniq command: %v", err)
    }

    headCmd := exec.Command("head", "-n", "10")
    headCmd.Stdin = strings.NewReader(uniqOut.String())
    var headOut bytes.Buffer
    headCmd.Stdout = &headOut

    err = headCmd.Run()
    if err != nil {
        t.Fatalf("Failed to run head command: %v", err)
    }

    lines := strings.Split(strings.TrimSpace(headOut.String()), "\n")

    // Check the first ten lines
    expectedLines := []string{
         "A",
        "ABACK",
        "ABANDON",
        "ABANDONED",
        "ABATED",
        "ABBREVIATED",
        "ABEYANCE",
        "ABIDE",
        "ABILITY",
        "ABLE",
    }

    if len(lines) != len(expectedLines) {
        t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(lines))
    }

    for i, line := range lines {
        if line != expectedLines[i] {
            t.Errorf("Line %d: expected '%s', got '%s'", i+1, expectedLines[i], line)
        }
    }

    // Print the first 10 lines for debugging
    t.Log("First 10 lines of output:")
    for _, line := range lines {
        t.Log(line)
    }
}

func TestMainWithInvalidArguments(t *testing.T) {
    // Save original os.Args
    oldArgs := os.Args
    defer func() { 
        os.Args = oldArgs 
    }()

    // Set up the new os.Args
    os.Args = []string{"cmd"}

    // Mock os.Exit
    oldStderr := os.Stderr
    r, w, _ := os.Pipe()
    os.Stderr = w
    defer func() {
        os.Stderr = oldStderr
    }()

    // Use a channel to communicate exit status
    exitChan := make(chan int, 1)

    oldExit := exit
    exit = func(code int) {
        exitChan <- code
    }
    defer func() {
        exit = oldExit
    }()

    // Run main in a goroutine
    go func() {
        main()
        close(exitChan)
    }()

    // Wait for exit or timeout
    var exitCode int
    select {
        case exitCode = <-exitChan:
        case <-time.After(5 * time.Second):
            t.Fatal("Test timed out")
    }

    // Close pipe and read output 
    w.Close()
    out, _ := io.ReadAll(r)

    t.Logf("Exit code: %d", exitCode)
    t.Logf("Stderr output: %q", string(out))

    if exitCode != 1 {
        t.Errorf("Expected exit code 1, but got %d", exitCode)
    }

    expected := "Usage: cmd <filename>\n"
    if string(out) != expected {
        t.Errorf("Expected output %q, but got %q", expected, string(out))
    }
}
