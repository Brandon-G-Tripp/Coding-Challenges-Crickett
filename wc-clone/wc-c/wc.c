#include <locale.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>
#include <wchar.h>

int count_chars(FILE* file) {
    setlocale(LC_ALL, "");

    int count = 0;
    wint_t ch;
    while ((ch = fgetwc(file)) != WEOF) {
        count++;
    } 

    return count;
}

int count_bytes(FILE* file) {
    int count = 0;
    int ch;
    while ((ch = fgetc(file)) != EOF) {
        count++;
    } 

    return (int)count;
} 

int count_lines(FILE* file) {
    int count = 0;
    char ch;
    int last_char = '\n';
    while ((ch = fgetc(file)) != EOF) {
        if (ch == '\n') {
            count++;
        } 
        last_char = ch;
    } 
    if (last_char != '\n' && last_char != EOF) {
        count++;
    } 

    return count;
}

int count_words(FILE* file) {
    int count = 0;
    int in_word = 0;
    char ch;
    while ((ch = fgetc(file)) != EOF) {
        if (ch == ' ' || ch == '\t' || ch == '\n') {
            if (in_word) {
                count++;
                in_word = 0;
            }
            in_word = 0;
        } else {
            in_word = 1;
        }
    } 
    if (in_word) {
        count++;
    }

    return count;
}
