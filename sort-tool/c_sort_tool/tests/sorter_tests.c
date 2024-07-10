#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include "../src/sorter.h"

#define TESTING

int test_sort_file();
int test_sort_main();
int test_sort_file_unique();
int test_sort_main_helper(int argc, char *argv[], const char *expected_output);

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

int test_sort_file_unique() {
    printf("Testing sort_file_unique function...\n");

    const char *filename = "test_unique_file.txt";
    FILE *file = fopen(filename, "w");
    if (file == NULL) {
        fprintf(stderr, "Error: Failed to create test file\n");
        fclose(file);
    }
    fprintf(file, "banana\napple\ncherry\nbanana\napple\n");
    fclose(file);

    int num_lines;
    char **lines = sort_file_unique(filename, &num_lines);
    if (lines == NULL) {
        fprintf(stderr, "Error: sort_file_unique failed\n");
        remove(filename);
        return 1;
    }

    char *expected_lines[] = {"apple\n", "banana\n", "cherry\n"};
    int expected_num_lines = 3;

    if (num_lines != expected_num_lines || !compare_lines(lines, expected_lines, num_lines)) {
        fprintf(stderr, "Test case for sort_file_unique\n");
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

    printf("All sort_file_unique tests passed.\n");
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
    fprintf(file, "ZEBRA\nABACK\nABANDON\nABATED\nABILITY\nABLE\nZEBRA\n");
    fclose(file);

    // Prepare arguments for sort_main test without -u option
    char *args[] = {"./sorter", (char *)filename};
    if (test_sort_main_helper(2, args, "ABACK\nABANDON\nABATED\nABILITY\nABLE\nZEBRA\nZEBRA\n") != 0) {
        remove(filename);
        return 1;
    }

    // Test with -u option
    char *args_u[] = {"./sorter", "-u", (char *)filename};
    if (test_sort_main_helper(3, args_u, "ABACK\nABANDON\nABATED\nABILITY\nABLE\nZEBRA\n") != 0) {
        remove(filename);
        return 1;
    }

    remove(filename);
    printf("All sort_main tests passed.\n");
    return 0;
}

int test_sort_main_helper(int argc, char *argv[], const char *expected_output) {
    FILE *temp_out = tmpfile();
    if (temp_out == NULL) {
        fprintf(stderr, "Error: Failed to create temporary file\n");
        return 1;
    }

    int old_stdout = dup(fileno(stdout));
    dup2(fileno(temp_out), fileno(stdout));

    int result = sort_file_from_args(argc, argv);

    fflush(stdout);
    dup2(old_stdout, fileno(stdout));
    close(old_stdout);

    if (result != 0) {
        fprintf(stderr, "Error: sort_main failed\n");
        fclose(temp_out);
        return 1;
    }

    rewind(temp_out);
    char buffer[1024];
    size_t expected_len = strlen(expected_output);
    size_t bytes_read = fread(buffer, 1, sizeof(buffer) - 1, temp_out);
    buffer[bytes_read] = '\0';

    if (bytes_read != expected_len || strncmp(buffer, expected_output, expected_len) != 0) {
        fprintf(stderr, "Test case for sort_main failed\n");
        fprintf(stderr, "Expected output:\n%s\n", expected_output);
        fprintf(stderr, "Actual output:\n%s\n", buffer);
        fclose(temp_out);
        return 1;
    }

    fclose(temp_out);
    return 0;
}

int main() {
    printf("Running tests... \n");

    if (test_sort_file() != 0 || test_sort_main() != 0 || test_sort_file_unique() != 0) {
        printf("Some tests failed.\n");
        return 1;
    }

    printf("All tests passed.\n");
    return 0;
}
