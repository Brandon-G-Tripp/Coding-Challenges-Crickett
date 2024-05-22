#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "cut.h"

bool cut_second_field(const char *file_path) {
    FILE *file = fopen(file_path, "r");
    if (file == NULL) {
        fprintf(stderr, "Failed to open file: %s\n", file_path);
        return false;
    }

    char line[1024];
    while (fgets(line, sizeof(line), file) != NULL) {
        char *token = strtok(line, "\t");
        if (token != NULL) {
            token = strtok(NULL, "\t");
            if (token != NULL) {
                printf("%s\n", token);
            }
        }
    }

    fclose(file);
    return true;
}
