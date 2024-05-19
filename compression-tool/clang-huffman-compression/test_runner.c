#include <stdio.h>

extern void test_count_character_frequency();
extern void test_count_character_frequency_empty_string();
extern void test_count_character_frequency_non_ascii();
extern void test_count_character_frequencies_from_file();

extern void test_create_leaf_node();
extern void test_create_internal_node();
extern void test_build_huffman_tree();

int main() {
   printf("Running frequency counter tests...\n");
   test_count_character_frequency();
   test_count_character_frequency_empty_string();
   test_count_character_frequency_non_ascii();
   test_count_character_frequencies_from_file();

   printf("\nRunning Huffman tree tests...\n");
   test_create_leaf_node();
   test_create_internal_node();
   test_build_huffman_tree();

   printf("\nAll tests passed!\n");
   return 0;
}
