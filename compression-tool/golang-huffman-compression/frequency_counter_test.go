package main

import "testing"

func TestCountCharacterFrequencies(t *testing.T) {
    input := "Hello, World!"
    expectedFrequencies := map[rune]int{
        'H': 1,
        'e': 1,
        'l': 3,
        'o': 2,
        ',': 1,
        ' ': 1,
        'W': 1,
        'r': 1,
        'd': 1,
        '!': 1,
    }

    actualFrequencies := CountCharacterFrequencies(input)

    for char, expectedCount := range expectedFrequencies {
        actualCount := actualFrequencies[char]
        if actualCount != expectedCount {
            t.Errorf("Expected %d occurrences of '%c', but got %d", expectedCount, char, actualCount)
        }
    }
}

