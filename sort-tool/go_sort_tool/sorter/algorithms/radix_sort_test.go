package algorithms

import (
    "reflect"
    "testing"
)

func TestRadixSort(t *testing.T) {
    tests := []struct {
        name string
        input []string
        expected []string
    }{
        {
            name:     "Empty slice",
            input:    []string{},
            expected: []string{},
        },
        {
            name:     "Single element",
            input:    []string{"AAA"},
            expected: []string{"AAA"},
        },
        {
            name:     "Already sorted",
            input:    []string{"AAA", "BBB", "CCC", "DDD"},
            expected: []string{"AAA", "BBB", "CCC", "DDD"},
        },
        {
            name:     "Reverse sorted",
            input:    []string{"DDD", "CCC", "BBB", "AAA"},
            expected: []string{"AAA", "BBB", "CCC", "DDD"},
        },
        {
            name:     "Random order",
            input:    []string{"CCC", "AAA", "DDD", "BBB"},
            expected: []string{"AAA", "BBB", "CCC", "DDD"},
        },
        {
            name:     "With duplicates",
            input:    []string{"BBB", "AAA", "CCC", "AAA", "BBB"},
            expected: []string{"AAA", "AAA", "BBB", "BBB", "CCC"},
        },
        {
            name:     "Different lengths",
            input:    []string{"A", "AA", "AAA", "AAAA"},
            expected: []string{"A", "AA", "AAA", "AAAA"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            RadixSort(tt.input)
            if !reflect.DeepEqual(tt.input, tt.expected) {
                t.Errorf("RadixSort() = %v, want %v", tt.input, tt.expected)
            }
        })
    }
}

