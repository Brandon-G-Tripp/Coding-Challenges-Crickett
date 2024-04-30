#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

extern int count_bytes(FILE* file);
extern int count_lines(FILE* file);
extern int count_words(FILE* file);
extern int count_chars(FILE* file);

void test_default_count() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Line 1\nLine 2\nLine 3\n");
    fclose(file);

    file = fopen("test.txt", "r");
    int line_count = count_lines(file);
    int word_count = count_words(file);
    int byte_count = count_bytes(file);
    int char_count = count_chars(file);
    fclose(file);

    assert(line_count == 3);
    assert(word_count == 6);
    assert(byte_count == 21);
    assert(char_count == 21);

    remove("test.txt");
} 

void test_count_chars() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Sample content with 🚀 emoji");
    fclose(file);

    file = fopen("test.txt", "r");
    int char_count = count_chars(file);
    int byte_count = count_bytes(file);
    fclose(file);

    assert(char_count == 27);
    assert(byte_count == 30);
    assert(char_count != byte_count);

    remove("test.txt");
} 

void test_count_bytes() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Sample content");
    fclose(file);

    // Call the count_bytes function
    int count = count_bytes(file);

    // Assert the expected byte count
    assert(count = 14);

    // clean up the temp file
    remove("test.txt");
}

void test_count_lines() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Line 1\nLine 2\nLine 3\n");
    fclose(file);

    file = fopen("test.txt", "r");
    int count = count_lines(file);
    fclose(file);

    assert(count == 3);

    remove("test.txt");
}

void test_count_words() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "This is a sample file\nwith multiple words\non each line\n");
    fclose(file);

    file = fopen("test.txt", "r");
    int count = count_words(file);
    fclose(file);

    assert(count == 11);

    remove("test.txt");
}

int main() {
    test_count_bytes();
    test_count_lines();
    test_count_words();
    test_count_chars();
    test_default_count();

    printf("All tests passed!\n");

    return 0;
} 
