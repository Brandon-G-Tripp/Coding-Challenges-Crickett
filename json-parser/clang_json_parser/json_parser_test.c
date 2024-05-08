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
    const char *json = "}";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON missing opening brace\n");
    } else {
        printf("Test passed: Invalid JSON missing opening brace\n");
    }
}

void test_valid_json_object_with_single_pair() {
    const char *json = "{\"key\":\"value\"}";
    bool is_valid = is_valid_json(json);
    if (is_valid) {
        printf("Test passed: Valid JSON object with single pair\n");
    } else {
        printf("Test failed: Valid JSON object with single pair\n");
    }
}

void test_valid_json_object_with_multiple_pairs() {
    const char *json = "{\"key1\":\"value1\",\"key2\":\"value2\"}";
    bool is_valid = is_valid_json(json);
    if (is_valid) {
        printf("Test passed: Valid JSON object with multiple pairs\n");
    } else {
        printf("Test failed: Valid JSON object with multiple pairs\n");
    }
}

void test_invalid_json_object_missing_key() {
    const char *json = "{:\"value\"}";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON object missing key\n");
    } else {
        printf("Test failed: Invalid JSON object missing key\n");
    }
}

void test_invalid_json_object_missing_value() {
    const char *json = "{\"key\":}";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON object missing value\n");
    } else {
        printf("Test failed: Invalid JSON object missing value\n");
    }
}

void test_invalid_json_object_missing_colon() {
    const char *json = "{\"key\"\"value\"}";
    bool is_valid = is_valid_json(json);
    if (!is_valid) {
        printf("Test passed: Invalid JSON object missing colon\n");
    } else {
        printf("Test failed: Invalid JSON object missing colon\n");
    }
}

void test_valid_json_object_with_different_value_types() {
    const char *json = "{\n"
            " \"key1\": true,\n"
            " \"key2\": false,\n"
            " \"key3\": null,\n"
            " \"key4\": \"value\",\n"
            " \"key5\": 101\n"
        "}";

    bool is_valid = is_valid_json(json);
    if (is_valid) {
        printf("Test passed: Valid JSON object with different value types\n");
    } else {
        printf("Test failed: Valid JSON object with different value types\n");
    }
}

int main() {
    test_valid_empty_object();
    test_valid_empty_object_with_whitespace();
    test_invalid_json_missing_closing_brace();
    test_invalid_json_missing_opening_brace();
    test_valid_json_object_with_single_pair();
    test_valid_json_object_with_multiple_pairs();
    test_invalid_json_object_missing_key();
    test_invalid_json_object_missing_value();
    test_invalid_json_object_missing_colon();
    test_valid_json_object_with_different_value_types();
    return 0;
} 
