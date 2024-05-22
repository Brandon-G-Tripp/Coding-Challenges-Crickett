#include <stdio.h>
#include <stdbool.h>
#include "cut.h"

int test_cut() {
    const char *file_path = "../sample.tsv";
    bool success = cut_second_field(file_path);

    if (success) {
        printf("Test passed\n");
        return 0;
    } else {
        printf("Test failed\n");
        return 1;
    }
}

