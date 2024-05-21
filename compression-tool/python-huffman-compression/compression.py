import sys

def count_frequencies(file):
    frequencies = {}
    for char in file.read():
        frequencies[char] = frequencies.get(char, 0) + 1
    return frequencies
