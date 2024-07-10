package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestMainWithInvalidArguments(t *testing.T) {
    testCases := []struct {
        name string
        args []string
        exitCode int
        output string
    }{
        {
            name:     "No arguments",
            args:     []string{"cmd"},
            exitCode: 1,
            output:   "Usage: cmd [-u] <filename>\n",
        },
        {
            name:     "-u flag without filename",
            args:     []string{"cmd", "-u"},
            exitCode: 1,
            output:   "Usage: cmd [-u] <filename>\n",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // reset flags
            flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
            // Save original os.Args and restore after test
            oldArgs := os.Args
            defer func() {
                os.Args = oldArgs
            }()

            // Set up the new os.Args
            os.Args = tc.args

            // Mock os.Exit and os.Stderr
            oldStderr := os.Stderr
            r, w, _ := os.Pipe()
            os.Stderr = w
            defer func() {
                os.Stderr = oldStderr
            }()

            exitChan := make(chan int, 1)
            oldExit := exit
            exit = func(code int) {
                exitChan <- code
            }
            defer func() {
                exit = oldExit
            }()

            // run main in a go routine
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

            if exitCode != tc.exitCode {
                t.Errorf("Expected exit code %d, but got %d", tc.exitCode, exitCode)
            }

            if string(out) != tc.output {
                t.Errorf("Expected output %q, but got %q", tc.output, string(out))
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
