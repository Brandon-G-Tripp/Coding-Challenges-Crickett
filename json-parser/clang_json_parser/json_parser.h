#ifndef JSON_PARSER_H
#define JSON_PARSER_H

#include <stdbool.h>

bool is_valid_json(const char *json);
void print_validation_result(bool is_valid);

#endif
