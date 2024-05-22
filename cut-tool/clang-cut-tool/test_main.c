#include <stdio.h>

int main() {
    printf("Running tests...\n");

    // Run the test functions
    extern int test_cut();
    int result = test_cut();

    if (result == 0) {
        printf("All tests passed!\n");
        return 0;
    } else {
        printf("Some tests failed.\n");
        return 1;
    }
}
