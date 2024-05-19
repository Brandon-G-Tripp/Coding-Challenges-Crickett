#include <stdio.h>
#include <stdlib.h>
#include <wchar.h>
#include "huffman_tree.h"

HuffmanNode* create_leaf_node(wchar_t character, int frequency) {
    HuffmanNode *node = (HuffmanNode*)malloc(sizeof(HuffmanNode));
    node->character = character;
    node->frequency = frequency;
    node->left = NULL;
    node->right = NULL;
    return node;
}

HuffmanNode* create_internal_node(HuffmanNode *left, HuffmanNode *right) {
    HuffmanNode *node = (HuffmanNode*)malloc(sizeof(HuffmanNode));
    node->character = L'\0';
    node->frequency = left->frequency + right->frequency;
    node->left = left;
    node->right = right;
    return node;
}

void free_huffman_tree(HuffmanNode *node) {
    if (node == NULL)
        return;

    free_huffman_tree(node->left);
    free_huffman_tree(node->right);
    free(node);
}

typedef struct MinHeapNode {
    HuffmanNode *node;
    struct MinHeapNode *next;
} MinHeapNode;

MinHeapNode* create_min_heap_node(HuffmanNode *node) {
    MinHeapNode *min_heap_node = (MinHeapNode*)malloc(sizeof(MinHeapNode));
    min_heap_node->node = node;
    min_heap_node->next = NULL;
    return min_heap_node;
}

void insert_min_heap(MinHeapNode **head, HuffmanNode *node) {
    MinHeapNode *new_node = create_min_heap_node(node);
    if (*head == NULL || (*head)->node->frequency >= node->frequency) {
        new_node->next = *head;
        *head = new_node;
    } else {
        MinHeapNode *current = *head;
        while (current->next != NULL && current->next->node->frequency < node->frequency) {
            current = current->next;
        }
        new_node->next = current->next;
        current->next = new_node;
    }
}

HuffmanNode* extract_min(MinHeapNode **head) {
    if (*head == NULL)
        return NULL;

    MinHeapNode *temp = *head;
    HuffmanNode *min_node = temp->node;
    *head = (*head)->next;
    free(temp);
    return min_node;
}

HuffmanNode* build_huffman_tree(int *character_frequencies) {
    MinHeapNode *head = NULL;

    printf("Building Huffman tree...\n");

    for (int i = 0; i < 256; i++) {
        if (character_frequencies[i] > 0) {
            printf("Creating leaf node for character '%c' with frequency %d\n", (char)i, character_frequencies[i]);
            HuffmanNode *leaf_node = create_leaf_node((wchar_t)i, character_frequencies[i]);
            printf("Inserting leaf node into min heap\n");
            insert_min_heap(&head, leaf_node);
        }
    }

    printf("Building internal nodes...\n");

    while (head != NULL && head->next != NULL) {
        printf("Extracting two minimum nodes from min heap\n");
        HuffmanNode *left = extract_min(&head);
        HuffmanNode *right = extract_min(&head);
        printf("Creating internal node with frequency %d\n", left->frequency + right->frequency);
        HuffmanNode *internal_node = create_internal_node(left, right);
        printf("Inserting internal node into min heap\n");
        insert_min_heap(&head, internal_node);
    }

    printf("Extracting root node from min heap\n");
    return extract_min(&head);
}


void traverse_and_print_huffman_tree(HuffmanNode *node, int depth) {
    if (node == NULL)
        return;

    for (int i = 0; i < depth; i++)
        printf(" ");

    if (node->character == L'\0')
        printf("Internal Node: Frequency = %d\n", node->frequency);
    else
        printf("Leaf Node: Character = '%c', Frequency = %d\n", node->character, node->frequency);

    traverse_and_print_huffman_tree(node->left, depth + 1);
    traverse_and_print_huffman_tree(node->right, depth + 1);
}
