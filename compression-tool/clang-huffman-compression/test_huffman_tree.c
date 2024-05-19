#include <stdio.h>
#include <stdlib.h>
#include <wchar.h>
#include "huffman_tree.h"

void test_create_leaf_node() {
    HuffmanNode *node = create_leaf_node('A', 5);
    if (node->character != 'A' || node->frequency != 5 || node->left != NULL || node->right !=NULL) {
        printf("Test failed: create_leaf_node\n");
        exit(1);
    }
    free(node);
    printf("Test passed: create_leaf_node\n");
}

void test_create_internal_node() {
    HuffmanNode *left = create_leaf_node('A', 5);
    HuffmanNode *right = create_leaf_node('B', 3);
    HuffmanNode *node = create_internal_node(left, right);
    if (node->character != '\0' || node->frequency != 8 || node->left != left || node->right != right) {
        printf("Test failed: create_internal_node\n");
        exit(1);
    }
    free_huffman_tree(node);
    printf("Test passed: create_internal_node\n");
}

void traverse_huffman_tree(HuffmanNode *node, wchar_t *expected_chars, int *expected_freqs, int num_expected, int *found_chars) {
    if (node == NULL)
        return;

    if (node->character != '\0') {
        for (int i = 0; i < num_expected; i++) {
            if (node->character == expected_chars[i] && node->frequency == expected_freqs[i]) {
                found_chars[node->character] = 1;
                break;
            }
        }
    }

    traverse_huffman_tree(node->left, expected_chars, expected_freqs, num_expected, found_chars);
    traverse_huffman_tree(node->right, expected_chars, expected_freqs, num_expected, found_chars);
}

void test_build_huffman_tree() {
    int character_frequencies[256] = {0};
    character_frequencies['A'] = 5;
    character_frequencies['B'] = 3;
    character_frequencies['C'] = 2;
    character_frequencies['D'] = 1;

    printf("Building huffman tree...\n");
    HuffmanNode *root = build_huffman_tree(character_frequencies);

    // Check if the root is NULL
    printf("Checking root node...\n");
    if (root == NULL) {
        printf("Assertion failed: Root is NULL\n");
        exit(1);
    }

    // Check the total frequency of the root node
    printf("Checking root frequency...\n");
    if (root->frequency != 11) {
        printf("Assertion failed: Root frequency is %d, expected 11\n", root->frequency);
        exit(1);
    }

    // Check if the left and right children of the root are not NULL
    printf("Checking left and right children of root...\n");
    if (root->left == NULL || root->right == NULL) {
        printf("Assertion failed: Left or right child of the root is NULL\n");
        exit(1);
    }

    // Check the frequencies of the left and right children of the root
    printf("Checking frequencies of left and right children...\n");
    if (root->left->frequency + root->right->frequency != 11) {
        printf("Assertion failed: Left child frequency is %d, expected 6. Right child frequency is %d, expected 5\n",
               root->left->frequency, root->right->frequency);
        exit(1);
    }

    wchar_t expected_chars[] = {'A', 'B', 'C', 'D'};
    int expected_freqs[] = {5, 3, 2, 1};
    int num_expected = sizeof(expected_chars) / sizeof(expected_chars[0]);
    int found_chars[65536] = {0};

    printf("Checking if all expected chars are found...\n");
    traverse_huffman_tree(root, expected_chars, expected_freqs, num_expected, found_chars);

    for (int i = 0; i < num_expected; i++) {
        if (!found_chars[expected_chars[i]]) {
            printf("Assertion failed: Character '%c' not found in Huffman tree\n", expected_chars[i]);
            exit(1);
        }
    }

    printf("Freeing Huffman tree...\n");
    free_huffman_tree(root);
    printf("Test passed: build_huffman_tree\n");
}

