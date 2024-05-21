#include <wchar.h>

void count_character_frequencies(wchar_t character, int *frequencies) {
    if (character == L' ') {
        frequencies[' ']++;
    } else {
        frequencies[character]++;
    }
}
