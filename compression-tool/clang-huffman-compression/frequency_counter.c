#include <wchar.h>
void count_character_frequencies(const wchar_t *input_string, int *frequencies) {
    while (*input_string != L'\0') {
        if (*input_string == L' ') {
            frequencies[L' ']++;
        } else {
            frequencies[(int)(*input_string)]++;
        }
        input_string++;
    }
}
