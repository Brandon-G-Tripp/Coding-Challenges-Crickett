#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "../src/sorter.h"

#define TESTING

int compare_lines(char **lines1, char **lines2, int num_lines) {
    for (int i = 0; i < num_lines; i++) {
        if (strcmp(lines1[i], lines2[i]) != 0) {
            return 0;
        }
    }
    return 1;
}

int test_sort_file() {
    printf("Testing sort_file function...\n");

    // Test case 1: Sort a small file
    const char *filename = "test_file.txt";
    FILE *file = fopen(filename, "w");
    if (file == NULL) {
        fprintf(stderr, "Error: Failed to create test file\n");
        return 1;
    }
    fprintf(file, "banana\napple\ncherry\n");
    fclose(file);

    int num_lines;
    char **lines = sort_file(filename, &num_lines);
    if (lines == NULL) {
        fprintf(stderr, "Error: sort_file failed\n");
        remove(filename);
        return 1;
    }

    char *expected_lines[] = {"apple\n", "banana\n", "cherry\n"};
    if (!compare_lines(lines, expected_lines, num_lines)) {
        fprintf(stderr, "Test case 1 failed\n");
        for (int i = 0; i < num_lines; i++) {
            free(lines[i]);
        }
        free(lines);
        remove(filename);
        return 1;
    }

    for (int i = 0; i < num_lines; i++) {
        free(lines[i]);
    }
    free(lines);
    remove(filename);

    printf("All sort_file tests passed.\n");
    return 0;

}

int test_sort_main() {
    printf("Testing sort_main function...\n");

    // Create a test file
    const char *filename = "test_main_file.txt";
    FILE *file = fopen(filename, "w");
    if (file == NULL) {
        fprintf(stderr, "Error: Failed to create test file\n");
        return 1;
    }
    fprintf(file, "ZEBRA\nABACK\nABANDON\nABATED\nABILITY\nABLE\n");
    fclose(file);

    // Prepare arguments for sort_main
    char *args[] = {"./sorter", (char *)filename};

    // redirect stdout to capture the output 
    FILE *temp_out = tmpfile();
    if (temp_out == NULL) {
        fprintf(stderr, "Error: Failed to create temporary file\n");
        remove(filename);
        return 1;
    }

    int old_stdout = dup(fileno(stdout));
    dup2(fileno(temp_out), fileno(stdout));

    // Run sort_main
    int result = sort_file_from_args(2, args);

    // Restore stdout
    fflush(stdout);
    dup2(old_stdout, fileno(stdout));
    close(old_stdout);

    if (result != 0) {
        fprintf(stderr, "Error: sort_main failed\n");
        remove(filename);
        fclose(temp_out);
        return 1;
    }

    // Check the output 
    rewind(temp_out);
    char buffer[1024];
    const char *expected_output = "ABACK\nABANDON\nABATED\nABILITY\nABLE\nZEBRA\n";
    size_t expected_len = strlen(expected_output);
    size_t bytes_read = fread(buffer, 1, sizeof(buffer) - 1, temp_out);
    buffer[bytes_read] = '\0';

    if (bytes_read != expected_len || strncmp(buffer, expected_output, expected_len) != 0) {
        fprintf(stderr, "Test case for sort_main failed\n");
        fprintf(stderr, "Expected output:\n%s\n", expected_output);
        fprintf(stderr, "Actual output:\n%s\n", buffer);
        remove(filename);
        fclose(temp_out);
        return 1;
    }

    fclose(temp_out);
    remove(filename);

    printf("All sort_main tests passed.\n");
    return 0;
}

int main() {
    printf("Running tests... \n");

    if (test_sort_file() != 0 || test_sort_main() != 0) {
        printf("Some tests failed.\n");
        return 1;
    }

    printf("All tests passed.\n");
    return 0;
}
