#ifndef SORTER_H
#define SORTER_H

char **sort_file(const char *filename, int *num_lines);
int sort_file_from_args(int argc, char *argv[]);
char **sort_file_unique(const char *filename, int *num_lines);

#endif
