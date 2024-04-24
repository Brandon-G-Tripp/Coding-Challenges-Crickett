#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

extern int count_bytes(const char* file_path);
extern int count_lines(const char* file_path);
extern int count_words(const char* file_path);

int main(int argc, char* argv[]) {
    int opt; 
    int count_bytes_flag = 0;
    int count_lines_flag = 0;
    int count_words_flag = 0;

    while ((opt = getopt(argc, argv, "clw")) != -1) {
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
            default:
                fprintf(stderr, "Usage: %s [-c] [-l] [-w] <file>\n", argv[0]);
                exit(EXIT_FAILURE);
        } 
    } 

    if (optind >= argc) {
        fprintf(stderr,"Missing file argument\n");
        fprintf(stderr, "Usage: %s [-c] [-l] [-w]<file>\n", argv[0]);
        exit(EXIT_FAILURE);
    } 

    const char* file_path = argv[optind];

    if (count_lines_flag) {
        int count = count_lines(file_path);
        if (count == -1) {
            fprintf(stderr, "Error: Could not open file '%s'\n", file_path);
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path);
    } 
    if (count_bytes_flag) {
        int count = count_lines(file_path);
        if (count == -1) {
            fprintf(stderr, "Error: Could not open file '%s'\n", file_path);
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path);
    }
    if (count_words_flag) {
        int count = count_words(file_path);
        if (count == -1) {
            fprintf(stderr, "Error: Could not open file '%s'\n", file_path);
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path);
    }

    if (!count_bytes_flag && !count_lines_flag && !count_words_flag) {
        fprintf(stderr, "Error: Missing -c, -l, or -w flag\n");
        fprintf(stderr, "Usage: %s [-c] [-l] [-w] <file\n", argv[0]);
        exit(EXIT_FAILURE);
    } 

    return 0;
}

