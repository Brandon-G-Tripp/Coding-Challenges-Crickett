#ifndef HUFFMAN_TREE_H
#define HUFFMAN_TREE_H

#include <stdbool.h>
#include <wchar.h>

typedef struct HuffmanNode {
    wchar_t character;
    int frequency;
    struct HuffmanNode *left;
    struct HuffmanNode *right;
} HuffmanNode;

HuffmanNode* create_leaf_node(wchar_t character, int frequency);
HuffmanNode* create_internal_node(HuffmanNode *left, HuffmanNode *right);
void free_huffman_tree(HuffmanNode *node);
HuffmanNode* build_huffman_tree(int *character_frequencies);
void traverse_and_print_huffman_tree(HuffmanNode *node, int depth);

#endif
