#include <stdio.h>
#include <stdlib.h>
#include <wchar.h>
#include "frequency_counter.h"

int main(int argc, char **argv) {
    if (argc < 2) {
        printf("Error: input file are missing\n");
        return -1;
    }
    FILE *file = fopen(argv[1], "r");

    if (file == NULL) {
        printf("Error: unable to open input file \n");
        return -1;
    }

    int character_frequencies[256] = {0};
    wchar_t character;

    while((character = fgetc(file)) != EOF) {
        count_character_frequencies(&character, character_frequencies);
    }

    fclose(file);
    return 0;
}
