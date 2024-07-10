#include "sorter.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char **sort_file(const char *filename, int *num_lines) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        fprintf(stderr, "Error: Failed to open file %s\n", filename);
        *num_lines = 0;
        return NULL;
    }

    // Count the number of lines
    int count = 0;
    char buffer[1024];
    while (fgets(buffer, sizeof(buffer), file) != NULL) {
        count++;
    }
    rewind(file);

    // Allocate memory for the lines
    char **lines = (char **)malloc(count * sizeof(char *));
    if (lines == NULL) {
        fprintf(stderr, "Error: Failed to allocate memory\n");
        fclose(file);
        *num_lines = 0;
        return NULL;
    }

    // Read the lines into the array
    int i = 0;
    while (fgets(buffer, sizeof(buffer), file) != NULL) {
        lines[i] = strdup(buffer);
        if (lines[i] == NULL) {
            fprintf(stderr, "Error: Failed to allocate memory \n");
            for (int j = 0; j < i; j++) {
                free(lines[j]);
            }
            free(lines);
            fclose(file);
            *num_lines = 0;
            return NULL;
        }
        i++;
    }

    fclose(file);
    *num_lines = count;

    // Sort the lines
    for (int i = 0; i < count - 1; i++) {
        for (int j = 0; j < count - i - 1; j++) {
            if (strcmp(lines[j], lines[j + 1]) > 0) {
                char *temp = lines[j];
                lines[j] = lines[j + 1];
                lines[j + 1] = temp;
            }
        }
    }

    return lines;
}

int sort_file_from_args(int argc, char *argv[]) {
    if (argc < 2 || argc > 3) {
        fprintf(stderr, "Usage: %s [-u] <filename>\n", argv[0]);
        return 1;
    }

    int unique = 0;
    const char *filename;

    if (argc == 3 && strcmp(argv[1], "-u") == 0) {
        unique = 1;
        filename = argv[2];
    } else {
        filename = argv[1];
    }
    
    // Call the sorting function here
    int num_lines;
    char **lines;

    if (unique) {
        lines = sort_file_unique(filename, &num_lines);
    } else {
        lines = sort_file(filename, &num_lines);
    }

    // Print the sorted lines
    for (int i = 0; i < num_lines; i++) {
        printf("%s", lines[i]);
        free(lines[i]);
    }
    free(lines);

    return 0;
}

char **sort_file_unique(const char *filename, int *num_lines) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        fprintf(stderr, "Error: Failed to open file %s\n", filename);
        *num_lines = 0;
        return NULL;
    }

    // Read lines into a dynamic array
    char **lines = NULL;
    int count = 0;
    char buffer[1024];
    lines = (char **)realloc(lines, (count + 1) * sizeof(char *));
    if (lines == NULL) {
        fprintf(stderr, "Error: Failed to allocate memory\n");
        fclose(file);
        *num_lines = 0;
        return NULL;
    }
    while(fgets(buffer, sizeof(buffer), file) != NULL) {
        // Check if the line already exists
        int is_unique = 1;
        for (int i = 0; i < count; i++) {
            if (strcmp(lines[i], buffer) == 0) {
                is_unique = 0;
                break;
            }
        }

        if (is_unique) {
            lines = (char **)realloc(lines, (count + 1) * sizeof(char *));
            if (lines == NULL) {
                fprintf(stderr, "Error: Failed to allocate memory\n");
                fclose(file);
                *num_lines = 0;
                return NULL;
            }
            lines[count] = strdup(buffer);
            if (lines[count] == NULL) {
                fprintf(stderr, "Error: Failed to allocate memory\n");
                for (int i = 0; i < count; i++) {
                    free(lines[i]);
                }
                free(lines);
                fclose(file);
                *num_lines = 0;
                return NULL;
            }
            count ++;
        }
    }

    fclose(file);
    *num_lines = count;

    // Sort lines
    for (int i = 0; i < count - 1; i++) {
        for (int j = 0; j < count - i - 1; j++) {
            if (strcmp(lines[j], lines[j + 1]) > 0) {
                char *temp = lines[j];
                lines[j] = lines[j + 1];
                lines[j + 1] = temp;
            }
        }
    }
    
    return lines;
}
