package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestMainWithInvalidArguments(t *testing.T) {
    testCases := []struct {
        name string
        args []string
        exitCode int
        outputPattern string
    }{
        {
            name:     "No arguments",
            args:     []string{},
            exitCode: 1,
            outputPattern: `(?s)^Error: Usage: .+`,
            // outputPattern: `(?s)^Error: Usage: .* [-u] [-a algorithm] <filename>\n$`,
        },
        {
            name:     "-u flag without filename",
            args:     []string{"-u"},
            exitCode: 1,
            outputPattern: `(?s)^Error: Usage: .+`,
            // outputPattern: `(?s)^Error: Usage: .* [-u] [-a algorithm] <filename>\n$`,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            cmd := exec.Command("go", append([]string{"run", "main.go"}, tc.args...)...)
            output, err := cmd.CombinedOutput()

            if err == nil {
                t.Errorf("Expected error, but got none")
            }

            exitError, ok := err.(*exec.ExitError)
            if !ok {
                t.Errorf("Expected exit error, but got: %v", err)
            } else if exitError.ExitCode() != tc.exitCode {
                t.Errorf("Expected exit code %d, but got %d", tc.exitCode, exitError.ExitCode())
            }

            matched, err := regexp.MatchString(tc.outputPattern, string(output))
            if err != nil {
                t.Fatalf("Error in regex matching: %v", err)
            }
            if !matched {
                t.Errorf("Output doesn't match expected pattern\nExpected pattern:%s\nGot: %s", tc.outputPattern, string(output))
            }
        })
    }
}

func TestCliWithWordsFile(t *testing.T) {
    dir, err := os.Getwd()
    if err != nil {
        t.Fatalf("Failed to get current working directory: %v", err)
    }

    tmpFile, err := ioutil.TempFile(dir, "words_*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tmpFile.Name())

    content := []byte("BANANA\nAPPLE\nCHERRY\nDATE\nBANANA\nAPPLE\n")
    if _, err := tmpFile.Write(content); err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
    if err := tmpFile.Close(); err != nil {
        t.Fatalf("Failed to close temp file: %v", err)
    }

    testCases := []struct {
        name string
        args []string
        expected []string
    }{
         {
            name: "Without -u flag",
            args: []string{filepath.Base(tmpFile.Name())},
            expected: []string{
                "APPLE", "APPLE", "BANANA", "BANANA", "CHERRY", "DATE",
            },
        },
        {
            name: "With -u flag",
            args: []string{"-u", filepath.Base(tmpFile.Name())},
            expected: []string{
                "APPLE", "BANANA", "CHERRY", "DATE",
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            cmd := exec.Command("go", append([]string{"run", "main.go"}, tc.args...)...)
            cmd.Dir = dir 
            output, err := cmd.CombinedOutput()
            if err != nil {
                t.Fatalf("Failed to run CLI command: %v\nOutput: %s", err, output)
            }

            lines := strings.Split(strings.TrimSpace(string(output)), "\n")
            if !reflect.DeepEqual(lines, tc.expected) {
                t.Errorf("Expected output:\n%s\nGot:\n%s", strings.Join(tc.expected, "\n"), strings.Join(lines, "\n"))
            }
        })
    }
}

func TestMainWithAlgorithmFlag(t *testing.T) {
    // Create a temporary file with some content
    content := "banana\napple\ncherry\ndate\n"
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())

    if _, err := tmpfile.Write([]byte(content)); err != nil {
        t.Fatal(err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    algorithms := []string{"merge", "quick", "heap", "radix"}
    expectedOutput := "APPLE\nBANANA\nCHERRY\nDATE\n"

    for _, algo := range algorithms {
        t.Run(algo, func(t *testing.T) {
            cmd := exec.Command("go", "run", "main.go", "-a", algo, tmpfile.Name())
            output, err := cmd.CombinedOutput()
            if err != nil {
                t.Fatalf("Failed to run CLI command: %v\nOutput: %s", err, output)
            }

            if string(output) != expectedOutput {
                t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, string(output))
            }
        })
    }
}

func TestMainWithInvalidAlgorithm(t *testing.T) {
    // Create a temporary file with some content
    content := "banana\napple\ncherry\ndate\n"
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())

    if _, err := tmpfile.Write([]byte(content)); err != nil {
        t.Fatal(err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    cmd := exec.Command("go", "run", "main.go", "-a", "invalid", tmpfile.Name())
    output, err := cmd.CombinedOutput()

    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    expectedWarning := "Warning: Unknown sorting algorithm 'invalid'. Defaulting to quick sort."
    expectedOutput := "APPLE\nBANANA\nCHERRY\nDATE\n"

    if !strings.Contains(string(output), expectedWarning) {
        t.Errorf("Expected error message to contain:\n%s\nGot:\n%s", expectedWarning, string(output))
    }

    if !strings.Contains(string(output), expectedOutput) {
        t.Errorf("Expected error message to contain:\n%s\nGot:\n%s", expectedOutput, string(output))
    }
}
