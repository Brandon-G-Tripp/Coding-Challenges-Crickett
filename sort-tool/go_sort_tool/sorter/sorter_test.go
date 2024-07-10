package sorter

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSortFile(t *testing.T) {
    content := "banana\napple\ncherry\napple\nbanana\n"
    filename := "test_input.txt"
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(filename)

    tests := []struct {
        name        string
        deduplicate bool
        algorithm   string
        expected    []string
    }{
        {
            name:        "No deduplication",
            deduplicate: false,
            algorithm:   "quick", // You can choose any default algorithm here
            expected:    []string{"APPLE", "APPLE", "BANANA", "BANANA", "CHERRY"},
        },
        {
            name:        "With deduplication",
            deduplicate: true,
            algorithm:   "quick", // You can choose any default algorithm here
            expected:    []string{"APPLE", "BANANA", "CHERRY"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := SortFile(filename, tt.deduplicate, tt.algorithm)
            if err != nil {
                t.Fatalf("SortFile failed: %v", err)
            }

            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("SortFile() = %v, want %v", result, tt.expected)
            }
        })
    }
}

func TestSortEmptyFile(t *testing.T) {
    filename := "empty.txt"
    err := ioutil.WriteFile(filename, []byte(""), 0644)
    if err != nil {
        t.Fatalf("Failed to create empty file: %v", err)
    }
    defer os.Remove(filename)

    result, err := SortFile(filename, false, "quick") // You can choose any default algorithm here
    if err != nil {
        t.Fatalf("SortFile failed: %v", err)
    }

    if len(result) != 0 {
        t.Errorf("Expected empty result, got %v", result)
    }
}

func TestSortFileWithMergeSort(t *testing.T) {
    content := "banana\napple\ncherry\ndate\n"
    filename := "test_merge_sort.txt"
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(filename)

    expected := []string{"APPLE", "BANANA", "CHERRY", "DATE"}
    result, err := SortFile(filename, false, "merge")
    if err != nil {
        t.Fatalf("SortFile failed: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("SortFile() with merge sort = %v, want %v", result, expected)
    }
}

func TestSortFileWithQuickSort(t *testing.T) {
    content := "banana\napple\ncherry\ndate\n"
    filename := "test_quick_sort.txt"
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(filename)

    expected := []string{"APPLE", "BANANA", "CHERRY", "DATE"}
    result, err := SortFile(filename, false, "quick")
    if err != nil {
        t.Fatalf("SortFile failed: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("SortFile() with quick sort = %v, want %v", result, expected)
    }
}

func TestSortFileWithHeapSort(t *testing.T) {
    content := "banana\napple\ncherry\ndate\n"
    filename := "test_heap_sort.txt"
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(filename)

    expected := []string{"APPLE", "BANANA", "CHERRY", "DATE"}
    result, err := SortFile(filename, false, "heap")
    if err != nil {
        t.Fatalf("SortFile failed: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("SortFile() with heap sort = %v, want %v", result, expected)
    }
}

func TestSortFileWithRadixSort(t *testing.T) {
    content := "banana\napple\ncherry\ndate\n"
    filename := "test_radix_sort.txt"
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(filename)

    expected := []string{"APPLE", "BANANA", "CHERRY", "DATE"}
    result, err := SortFile(filename, false, "radix")
    if err != nil {
        t.Fatalf("SortFile failed: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("SortFile() with radix sort = %v, want %v", result, expected)
    }


}
