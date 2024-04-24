#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int count_bytes(const char* file_path) {
    FILE* file = fopen(file_path, "rb");
    if (file == NULL) {
        return -1; // this indicates error
    } 

    fseek(file, 0, SEEK_END);
    long count = ftell(file);
    fclose(file);

    return (int)count;
} 

int count_lines(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (file == NULL) {
        return -1;
    } 

    int count = 0;
    char ch;
    while ((ch = fgetc(file)) != EOF) {
        if (ch == '\n') {
            count++;
        } 
    } 
    fclose(file);

    return count;
}

int count_words(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (file == NULL) {
        return -1;
    }

    int count = 0;
    int in_word = 0;
    char ch;
    while ((ch = fgetc(file)) != EOF) {
        if (ch == ' ' || ch == '\t' || ch == '\n') {
            if (in_word) {
                count++;
                in_word = 0;
            }
        } else {
            in_word = 1;
        }
    } 
    if (in_word) {
        count++;
    }
    fclose(file);

    return count;
}
