#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "cut.h"

int main(int argc, char *argv[]) {
    if (argc != 3 || strcmp(argv[1], "-f2") != 0) {
        fprintf(stderr, "Usage: cut -f2 <file\n");
        return 1;
    }

    const char *file_path = argv[2];
    bool success = cut_second_field(file_path);

    if (!success) {
        return 1;
    }

    return 0;
}
