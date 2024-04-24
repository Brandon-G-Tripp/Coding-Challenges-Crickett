#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

extern int count_bytes(const char* file_path);
extern int count_lines(const char* file_path);
extern int count_words(const char* file_path);
extern int count_chars(const char* file_path);

void test_count_chars() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Sample content with 🚀 emoji");
    fclose(file);

    int char_count = count_chars("test.txt");
    int byte_count = count_bytes("test.txt");

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
    int count = count_bytes("test.txt");

    // Assert the expected byte count
    assert(count = 14);

    // clean up the temp file
    remove("test.txt");
}

void test_file_not_found() {
    //' Call the count_+bytes function with a non-existent file
    int count = count_bytes("nonexistent.txt");

    // Assert that the coutn is -1 (indicating error)
    assert(count == -1);
} 

void test_count_lines() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "Line 1\nLine 2\nLine 3\n");
    fclose(file);

    int count = count_lines("test.txt");

    assert(count == 3);

    remove("test.txt");
}

void test_count_words() {
    FILE* file = fopen("test.txt", "w");
    fprintf(file, "This is a sample file\nwith multiple words\non each line\n");
    fclose(file);

    int count = count_words("test.txt");

    assert(count == 11);

    remove("test.txt");
}

int main() {
    test_count_bytes();
    test_file_not_found();
    test_count_lines();
    test_count_words();
    test_count_chars();

    printf("All tests passed!\n");

    return 0;
} 
