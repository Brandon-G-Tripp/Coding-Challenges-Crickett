#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <ctype.h>
#include "json_parser.h"

bool is_valid_value(const char *value) {
    const char *ptr = value;

    while (isspace(*ptr)) {
        ptr++;
    }

    if (strcmp(ptr, "true") == 0 || strcmp(ptr, "false") == 0 || strcmp(ptr, "null") == 0) {
        return true;
    }

    if (*ptr == '"') {
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
        return *ptr == '\0';
    }

    if (*ptr == '-') {
        ptr++;
    }

    if (!isdigit(*ptr)) {
        return false;
    }

    while (isdigit(*ptr)) {
        ptr++;
    }

    if (*ptr == '.') {
        ptr++;
        if (!isdigit(*ptr)) {
            return false;
        }
        while (isdigit(*ptr)) {
            ptr++;
        }
    }

    while (isspace(*ptr)) {
        ptr++;
    }

    return *ptr == '\0';
}

bool is_valid_json(const char *json) {
    const char *ptr = json;

    while (isspace(*ptr)) {
        ptr++;
    }

    if (*ptr != '{') {
        return false;
    }
    ptr++;

    while (isspace(*ptr)) {
        ptr++;
    }

    if (*ptr == '}') {
        return true;
    }

    while (*ptr != '\0') {
        if (*ptr != '"') {
            return false;
        }
        ptr++;
        const char *key_start = ptr;
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

        const char *value_start = ptr;
        while (*ptr != ',' && *ptr != '}' && *ptr != '\0') {  // Added a check for end of string
            ptr++;
        }

        char value[256];
        strncpy(value, value_start, ptr - value_start);
        value[ptr - value_start] = '\0';

        if (!is_valid_value(value)) {
            return false;
        }

        while (isspace(*ptr)) {
            ptr++;
        }

        if (*ptr == ',') {
            ptr++;
            while (isspace(*ptr)) {
                ptr++;
            }
            if (*ptr == '}') {  // Added a check for trailing comma
                return false;
            }
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
