#include <stdio.h>
#include <stdbool.h>
#include"json_parser.h"

void test_valid_empty_object() {
    const char *json = "{}";
    bool is_valid = is_valid_json(json);
    if (is_valid) {
        printf("Test passed: Valid empty object\n");
    } else {
        printf("Test failed: Valid empty object\n");
    }
}

void test_valid_empty_object_with_whitespace() {
    const char *json = " \t\n{}\r\n ";
    bool is_valid = is_valid_json(json);

    if (is_valid) {
        printf("Test passed: Valid empty object with whitespace\n");
    } else {
        printf("Test failed: Valid empty object with whitespace\n");
    }
} 

void test_invalid_json_missing_closing_brace() {
    const char *json = "{";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON missing closing brace\n");
    } else {
        printf("Test passed: Invalid JSON missing closing brace\n");
    }
}

void test_invalid_json_missing_opening_brace() {
    const char *json = "{";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON missing opening brace\n");
    } else {
        printf("Test passed: Invalid JSON missing opening brace\n");
    }
}

int main() {
    test_valid_empty_object();
    test_valid_empty_object_with_whitespace();
    test_invalid_json_missing_closing_brace();
    test_invalid_json_missing_opening_brace();
    return 0;
} 
