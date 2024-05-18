package main

func CountCharacterFrequencies(input string) map[rune]int {
    frequencies := make(map[rune]int)
    for _, char := range input {
        frequencies[char]++
    }
    return frequencies
}
