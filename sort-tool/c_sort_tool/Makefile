# Compiler
CC = gcc

# Compiler flags
CFLAGS = -Wall -Wextra -g

# Source files
SRC_FILES = $(wildcard src/*.c)
TEST_FILES = $(wildcard tests/*.c)

# Object files
OBJ_FILES = $(patsubst %.c, %.o, $(SRC_FILES))
TEST_OBJ_FILES = $(patsubst %.c, %.o, $(TEST_FILES))

# Executable
EXECUTABLE = sorter

# Test executable 
TEST_EXECUTABLE = test_sorter

# Default target
all: $(EXECUTABLE) $(TEST_EXECUTABLE)

# Build the executable 
$(EXECUTABLE): $(OBJ_FILES)
	$(CC) $(CFLAGS) $(filter src/main.o, $^) $(filter src/sorter.o, $^) -o $@

# Build the test executable
$(TEST_EXECUTABLE): $(filter-out src/main.o, $(OBJ_FILES)) $(TEST_OBJ_FILES)
	$(CC) $(CFLAGS) $^ -o $@

# Clean up 
clean:
	rm -f $(EXECUTABLE) $(TEST_EXECUTABLE) $(OBJ_FILES) $(TEST_OBJ_FILES)

# Compile source files
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

test: $(TEST_EXECUTABLE)
	./$(TEST_EXECUTABLE)

.PHONY: all clean test
