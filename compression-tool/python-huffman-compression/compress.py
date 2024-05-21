import sys
from compression import count_frequencies
from huffman_tree import build_huffman_tree, HuffmanNode

def print_huffman_tree(node, prefix=''):
    if node.char is not None:
        print(f"{prefix}Leaf: {node.char} (Frequency: {node.freq})")
    else:
        print(f"{prefix}Internal Node (Frequency: {node.freq})")
        print_huffman_tree(node.left, prefix + ' ')
        print_huffman_tree(node.right, prefix + ' ')

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python compress.py <filename>")
        sys.exit(1)

    filename = sys.argv[1]
    try:
        with open(filename, 'r') as file: 
            frequencies = count_frequencies(file)
            print("Character Frequencies:")
            print(frequencies)

            huffman_tree = build_huffman_tree(frequencies)
            print("Huffman Tree:")
            print_huffman_tree(huffman_tree)
    except FileNotFoundError:
        print("Error: File not found")
        sys.exit(1)


