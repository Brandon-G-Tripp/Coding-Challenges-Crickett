import sys

def count_frequencies(file):
    frequencies = {}
    for char in file.read():
        frequencies[char] = frequencies.get(char, 0) + 1
    return frequencies

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python compression.py <filename>")
        sys.exit(1)

    filename = sys.argv[1]
    try:
        with open(filename, 'r') as file:
            frequencies = count_frequencies(file)
            print(frequencies)
    except FileNotFoundError:
        print("Error: File not found")
        sys.exit(1)
