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

int main(int argc, char* argv[]) {
    int opt; 
    int count_bytes_flag = 0;

    while ((opt = getopt(argc, argv, "c")) != -1) {
        switch (opt) {
            case 'c':
                count_bytes_flag = 1;
                break;
            default:
                fprintf(stderr, "Usage: %s [-c] <file>\n", argv[0]);
                exit(EXIT_FAILURE);
        } 
    } 

    if (optind >= argc) {
        fprintf(stderr,"Missing file argument\n");
        fprintf(stderr, "Usage: %s [-c] <file>\n", argv[0]);
        exit(EXIT_FAILURE);
    } 

    const char* file_path = argv[optind];

    if (count_bytes_flag) {
        int count = count_bytes(file_path);
        if (count == -1) {
            fprintf(stderr, "Error: Could not open file '%s'\n", file_path);
            exit(EXIT_FAILURE);
        } 
        printf("%d %s\n", count, file_path);
    } else {
        fprintf(stderr, "Error: Missing -c flag\n");
        fprintf(stderr, "Usage: %s [-c] <file>\n", argv[0]);
        exit(EXIT_FAILURE);
    }

    return 0;
}

