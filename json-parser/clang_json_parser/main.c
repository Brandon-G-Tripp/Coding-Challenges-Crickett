#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "json_parser.h"

#define MAX_JSON_SIZE 1024

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <json_file>\n", argv[0]);
        exit(1);
    } 

    const char *filename = argv[1];
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        fprintf(stderr, "Error opening file: %s\n", filename);
        exit(1);
    } 

    char json[MAX_JSON_SIZE];
    size_t len = fread(json, 1, sizeof(json) - 1, file);
    json[len] = '\0';
    fclose(file);

    if (len > 0 && json[len - 1] == '\n') {
        json[len - 1] = '\0';
    } 

    bool is_valid = is_valid_json(json);
    print_validation_result(is_valid);

    return is_valid ? 0 : 1;
}
