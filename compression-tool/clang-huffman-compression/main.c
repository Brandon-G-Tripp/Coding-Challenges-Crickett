#include <stdio.h>
#include <stdlib.h>
#include <wchar.h>
#include "frequency_counter.h"
#include "huffman_tree.h"

int main(int argc, char **argv) {
    printf("Huffman Compression\n");

    if (argc < 2) {
        printf("Error: input file is missing\n");
        return -1;
    }

    printf("Opening input file: %s\n", argv[1]);
    FILE *file = fopen(argv[1], "r");

    if (file == NULL) {
        printf("Error: unable to open input file\n");
        return -1;
    }

    printf("Allocating character frequencies array\n");
    int *character_frequencies = (int *)calloc(65536, sizeof(int));
    wchar_t character;

    printf("Counting character frequencies\n");
    while ((character = fgetwc(file)) != WEOF) {
        count_character_frequencies(character, character_frequencies);
    }

    printf("Closing input file\n");
    fclose(file);

    printf("Character frequencies: ");
    for (int i = 0; i < 65536; i++) {
        if (character_frequencies[i] > 0) {
            printf("'%lc': %d, ", (wchar_t)i, character_frequencies[i]);
        }
    }
    printf("\n");

    printf("Building Huffman tree\n");
    HuffmanNode *root = build_huffman_tree(character_frequencies);

    if (root == NULL) {
        printf("Error: Failed to build Huffman tree\n");
        free(character_frequencies);
        return -1;
    }

    printf("Traversing and printing Huffman tree\n");
    traverse_and_print_huffman_tree(root, 0);

    printf("Freeing Huffman tree\n");
    free_huffman_tree(root);

    printf("Freeing character frequencies array\n");
    free(character_frequencies);

    printf("Huffman compression completed successfully\n");
    return 0;
}
