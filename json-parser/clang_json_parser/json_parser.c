#include <stdio.h>
#include <string.h>
#include "json_parser.h"

bool is_valid_json(const char *json) {
    size_t len = strlen(json);
    size_t i = 0;

    while (i < len && (json[i] == ' ' || json[i] == '\t' || json[i] == '\n' || json[i] == '\r')) {
        i++;
    }

    if (i >= len || json[i] != '{') {
        return false;
    } 

    while (len > 0 && (json[len - 1] == ' ' || json[len - 1] == '\t' || json[len - 1] == '\n' || json[len - 1] == '\r')) {
        len--;
    } 

    if (len == 0 || json[len - 1] != '}') {
        return false;
    } 

    return true;
} 

void print_validation_result(bool is_valid) {
    if (is_valid) {
        printf("Valid JSON\n");
    } else {
        printf("Invalid JSON\n");
    }
} 
