package algorithms

import (
    "reflect"
    "testing"
)

func TestMergeSort(t *testing.T) {
    tests := []struct {
        name string          
        input []string
        expected []string
    } {
        {
            name:     "Empty slice",
            input:    []string{},
            expected: []string{},
        },
        {
            name:     "Single element",
            input:    []string{"A"},
            expected: []string{"A"},
        },
        {
            name:     "Already sorted",
            input:    []string{"A", "B", "C", "D"},
            expected: []string{"A", "B", "C", "D"},
        },
        {
            name:     "Reverse sorted",
            input:    []string{"D", "C", "B", "A"},
            expected: []string{"A", "B", "C", "D"},
        },
        {
            name:     "Random order",
            input:    []string{"B", "D", "A", "C"},
            expected: []string{"A", "B", "C", "D"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            MergeSort(tt.input)
            if !reflect.DeepEqual(tt.input, tt.expected) {
                t.Errorf("MergeSort() = %v, want %v", tt.input, tt.expected)
            }
        })
    }
}
