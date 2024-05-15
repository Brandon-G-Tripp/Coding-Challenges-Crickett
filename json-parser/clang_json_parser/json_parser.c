#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <stdlib.h>
#include <ctype.h>
#include "json_parser.h"

#define MAX_STACK_SIZE 256

char stack[MAX_STACK_SIZE];
int top = -1;

void push(char c) {
    if (top == MAX_STACK_SIZE - 1) {
        // stack overflow
        return;
    }
    stack[++top] = c;
}

char pop() {
    if (top == -1) {
        // stack underflow
        return '\0';
    }
    return stack[top--];
}

bool is_valid_value(const char **ptr);

bool is_valid_string(const char **ptr) {
    const char *p = *ptr;

    if (*p != '"') {
        return false;
    }
    p++;

    while (*p != '"') {
        if (*p == '\0') {
            return false;
        }
        p++;
    }
    p++;

    *ptr = p;
    return true;
}

bool is_valid_number(const char **ptr) {
    const char *p = *ptr;

    if (*p == '-') {
        p++;
    }

    if (!isdigit(*p)) {
        return false;
    }

    while (isdigit(*p)) {
        p++;
    }

    if (*p == '.') {
        p++;
        if (!isdigit(*p)) {
            return false;
        }
        while (isdigit(*p)) {
            p++;
        }
    }

    *ptr = p;
    return true;
}

bool is_valid_object(const char **ptr) {
    const char *p = *ptr;

    if (*p != '{') {
        return false;
    }
    p++;

    while (isspace(*p)) {
        p++;
    }

    if (*p == '}') {
        p++;
        *ptr = p;
        return true;
    }

    while (*p != '\0') {
        if (!is_valid_string(&p)) {
            return false;
        }

        while (isspace(*p)) {
            p++;
        }

        if (*p != ':') {
            return false;
        }
        p++;

        while (isspace(*p)) {
            p++;
        }

        if (!is_valid_value(&p)) {
            return false;
        }

        while (isspace(*p)) {
            p++;
        }

        if (*p == ',') {
            p++;
            while (isspace(*p)) {
                p++;
            }
        } else if (*p == '}') {
            p++;
            break;
        } else {
            return false;
        }
    }

    *ptr = p;
    return true;
}

bool is_valid_array(const char **ptr) {
    const char *p = *ptr;

    if (*p != '[') {
        return false;
    }
    p++;

    while (isspace(*p)) {
        p++;
    }

    if (*p == ']') {
        p++;
        *ptr = p;
        return true;
    }

    while (*p != '\0') {
        if (!is_valid_value(&p)) {
            return false;
        }

        while (isspace(*p)) {
            p++;
        }

        if (*p == ',') {
            p++;
            while (isspace(*p)) {
                p++;
            }
        } else if (*p == ']') {
            p++;
            break;
        } else {
            return false;
        }
    }

    *ptr = p;
    return true;
}

bool is_valid_value(const char **ptr) {
    const char *p = *ptr;

    while (isspace(*p)) {
        p++;
    }

    if (*p == '"') {
        if (!is_valid_string(&p)) {
            return false;
        }
    } else if (*p == '-' || isdigit(*p)) {
        if (!is_valid_number(&p)) {
            return false;
        }
    } else if (*p == '{') {
        if (!is_valid_object(&p)) {
            return false;
        }
    } else if (*p == '[') {
        if (!is_valid_array(&p)) {
            return false;
        }
    } else if (strncmp(p, "true", 4) == 0) {
        p += 4;
    } else if (strncmp(p, "false", 5) == 0) {
        p += 5;
    } else if (strncmp(p, "null", 4) == 0) {
        p += 4;
    } else {
        return false;
    }

    *ptr = p;
    return true;
}

bool is_valid_json(const char *json) {
    const char *ptr = json;

    while (isspace(*ptr)) {
        ptr++;
    }

    if (*ptr == '{') {
        if (!is_valid_object(&ptr)) {
            return false;
        }
    } else if (*ptr == '[') {
        if (!is_valid_array(&ptr)) {
            return false;
        }
    } else {
        return false;
    }

    while (isspace(*ptr)) {
        ptr++;
    }

    return *ptr == '\0';
}

void print_validation_result(bool is_valid) {
    if (is_valid) {
        printf("Valid JSON\n");
    } else {
        printf("Invalid JSON\n");
    }
}
