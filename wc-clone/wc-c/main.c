#include <locale.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

extern int count_bytes(FILE* file);
extern int count_lines(FILE* file);
extern int count_words(FILE* file);
extern int count_chars(FILE* file);

int main(int argc, char* argv[]) {
    setlocale(LC_ALL,"");
    int opt; 
    int count_bytes_flag = 0;
    int count_lines_flag = 0;
    int count_words_flag = 0;
    int count_chars_flag = 0;

    while ((opt = getopt(argc, argv, "clwm")) != -1) {
        switch (opt) {
            case 'c':
                count_bytes_flag = 1;
                break;
            case 'l':
                count_lines_flag = 1;
                break;
            case 'w':
                count_words_flag = 1;
                break;
            case 'm':
                count_chars_flag = 1;
                break;
            default:
                fprintf(stderr, "Usage: %s [-c] [-l] [-w] [-m] <file>\n", argv[0]);
                exit(EXIT_FAILURE);
        } 
    } 

    FILE* file = stdin;
    const char* file_path = NULL;


    if (optind < argc) {
        file_path = argv[optind];
        file = fopen(file_path, "r");
        if (file == NULL) {
            fprintf(stderr, "Error: Could not open file '%s'\n", file_path);
            exit(EXIT_FAILURE);
        } 
    } 

    if (count_chars_flag) {
        int count = count_chars(file);
        if (count == -1) {
            fprintf(stderr, "Error: Could not process input\n");
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path ? file_path : "");
    } 

    if (count_lines_flag) {
        int count = count_lines(file);
        if (count == -1) {
            fprintf(stderr, "Error: Could not process input\n");
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path ? file_path : "");
    } 
    if (count_bytes_flag) {
        int count = count_bytes(file);
        if (count == -1) {
            fprintf(stderr, "Error: Could not process input\n");
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path ? file_path : "");
    }
    if (count_words_flag) {
        int count = count_words(file);
        if (count == -1) {
            fprintf(stderr, "Error: Could not process input\n");
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path ? file_path : "");
    }

    if (!count_bytes_flag && !count_lines_flag && !count_words_flag && !count_chars_flag) {
        int line_count = 0;
        int word_count = 0;
        int byte_count = 0;
        int char_count = 0;
        int in_word = 0;
        char ch;

        while ((ch = fgetc(file)) != EOF) {
            byte_count++;

            if (ch == 'n') {
                line_count++;
            }

            if (ch == ' ' || ch == '\t' || ch == '\n') {
                if (in_word) {
                    word_count++;
                } 
                in_word = 0;
            } else {
                in_word = 1;
            }

            char_count++;
        } 

        if (in_word) {
            word_count++;
        }

        printf("%d %d %d %d\n", line_count, word_count, char_count, byte_count);
        return 0;
    } 

    if (file != stdin) {
        fclose(file);
    } 

    return 0;
}

