#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <ctype.h>
#include "json_parser.h"

bool is_valid_json(const char *json) {
    const char *ptr = json;

    while (isspace(*ptr)) {
        ptr++;
    } 

    if (*ptr != '{') {
        return false;
    }
    ptr++;

    while(isspace(*ptr)) {
        ptr++;
    } 

    while (*ptr != '\0') {
        if (*ptr != '"') {
            return false;
        }
        ptr++;
        while (*ptr != '"') {
            if (*ptr == '\0') {
                return false;
            }
            ptr++;
        }
        ptr++;

        while (isspace(*ptr)) {
            ptr++;
        }

        if (*ptr != ':') {
            return false;
        }
        ptr++;

        while (isspace(*ptr)) {
            ptr++;
        }

        if (*ptr != '"') {
            return false;
        }
        ptr++;
        while (*ptr != '"') {
            if (*ptr == '\0') {
                return false;
            }
            ptr++;
        }
        ptr++;

        while (isspace(*ptr)) {
            ptr++;
        }

        if (*ptr == ',') {
            ptr++;
        } else if (*ptr == '}') {
            return true;
        } else {
            return false;
        }
    } 

    return false;
} 

void print_validation_result(bool is_valid) {
    if (is_valid) {
        printf("Valid JSON\n");
    } else {
        printf("Invalid JSON\n");
    }
} 
