package cut

import (
    "io/ioutil"
    "os"
    "testing"
)

func TestCutSecondField(t *testing.T) {
    // Create temp file with sample data
    content := []byte("f0\tf1\tf2\tf3\tf4\tf5\n0\t1\t2\t3\t4\t5\n5\t6\t7\t8\t9\t10\n")
    tmpfile, err := ioutil.TempFile("", "sample")
    if err != nil {
        t.Fatalf("failed to create temporary file: %v", err)
    }
    defer os.Remove(tmpfile.Name())

    if _, err := tmpfile.Write(content); err != nil {
        t.Fatalf("failed to write to temporary file: %v", err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatalf("failed to close temporary file: %v", err)
    }

    // Call the CutSecondField function
    if err := CutSecondField(tmpfile.Name()); err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Check the output
    expectedOutput := "f1\n1\n6\n"
    if output := captureOutput(func() { CutSecondField(tmpfile.Name()) }); output != expectedOutput {
        t.Errorf("expected output %q, but got %q", expectedOutput, output)
    }
}

func captureOutput(f func()) string{
    rescueStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
    
    f()

    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = rescueStdout

    return string(out)
}
