#include <stdio.h>
#include <stdlib.h>
#include <wchar.h>
#include "frequency_counter.h"

#define BUFFER_SIZE 1024
#define TEST_FILE "../LesMiserables.txt"

void count_character_frequencies(const wchar_t *input_string, int *frequencies);

void test_count_character_frequency() {
    const wchar_t *input_string = L"Hello, World!";
    wchar_t expected_frequencies[256] = {0};
    expected_frequencies['H'] = 1;
    expected_frequencies['e'] = 1;
    expected_frequencies['l'] = 3;
    expected_frequencies['o'] = 2;
    expected_frequencies[','] = 1;
    expected_frequencies['W'] = 1;
    expected_frequencies['r'] = 1;
    expected_frequencies['d'] = 1;
    expected_frequencies['!'] = 1;
    expected_frequencies[' '] = 1;

    wchar_t actual_frequencies[256] = {0};
    count_character_frequencies(input_string, actual_frequencies);

    for (int i = 0; i < 256; i++) {
        if (expected_frequencies[i] != actual_frequencies[i]) {
            printf("Test failed - test_count_character_frequency: Expected %d occurrences of '%c', but got %d\n",
                    expected_frequencies[i], i, actual_frequencies[i]);
            return;
        }
    }


    printf("Test passed - test_count_character_frequency \n");
}

void test_count_character_frequency_empty_string() {
    const wchar_t *input_string = L"";
    wchar_t frequencies[256] = {0};
    count_character_frequencies(input_string, frequencies);

    for (int i = 0; i < 256; i++) {
        if (frequencies[i] != 0) {
            printf("Test failed - test_count_character_frequency_empty_string: Expected 0 occurrences of all characters, but got non-zero for '%c'\n", i);
            return;
        }
    }

    printf("Test passed  test_count_character_frequency_empty_string\n");
}

void test_count_character_frequency_non_ascii() {
    const wchar_t *input_string = L"Hëllö, Wörld!";
    wchar_t expected_frequencies[256] = {0};
    expected_frequencies[L'H'] = 1;
    expected_frequencies[L'ë'] = 1;
    expected_frequencies[L'l'] = 3; // 'l' appears twice
    expected_frequencies[L'ö'] = 2;
    expected_frequencies[L','] = 1;
    expected_frequencies[L' '] = 1;
    expected_frequencies[L'W'] = 1;
    expected_frequencies[L'r'] = 1;
    expected_frequencies[L'd'] = 1;
    expected_frequencies[L'!'] = 1;

    wchar_t actual_frequencies[256] = {0};
    count_character_frequencies(input_string, actual_frequencies);

    for (int i = 0; i < 256; i++) {
        if (expected_frequencies[i] != actual_frequencies[i]) {
            printf("Test failed - test_count_character_frequency_non_ascii: Expected %d occurrences of '%c', but got %d\n",
                   expected_frequencies[i], i, actual_frequencies[i]);
            return;
        }
    }

    printf("Test passed test_count_character_frequency_non_ascii\n");
}

void test_count_character_frequencies_from_file() {
    FILE *file = fopen(TEST_FILE, "r");
    if (file == NULL) {
        printf("Error: Unable to open test file %s\n", TEST_FILE);
        exit(1);
    }

    wchar_t buffer[BUFFER_SIZE];
    wchar_t character_frequencies[256] = {0};

    while (fgetws(buffer, BUFFER_SIZE, file)) {
        count_character_frequencies(buffer, character_frequencies);
    }

    fclose(file);

    if (character_frequencies['X'] != 333) {
        printf("Test failed - test_count_character_frequencies_from_file: Expected 333 occurrences of 'X', but got %d\n", character_frequencies['X']);
        exit(1);
    }

    if (character_frequencies['t'] != 223000) {
        printf("Test failed test_count_character_frequency_non_ascii: Expected 223000 occurrences of 't', but got %d\n", character_frequencies['t']);
        exit(1);
    }

    printf("Test passed test_count_character_frequency_non_ascii\n");
}

