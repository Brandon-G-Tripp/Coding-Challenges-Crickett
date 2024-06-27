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
        lines[i] = (char *)malloc((strlen(buffer) + 1) * sizeof(char));
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
        strcpy(lines[i], buffer);
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
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <filename>\n", argv[0]);
        return 1;
    }
    
    // Call the sorting function here
    int num_lines;
    char **lines = sort_file(argv[1], &num_lines);
    if (lines == NULL) {
        return 1;
    }

    // Print the sorted lines
    for (int i = 0; i < num_lines; i++) {
        printf("%s", lines[i]);
        free(lines[i]);
    }
    free(lines);

    return 0;
}
