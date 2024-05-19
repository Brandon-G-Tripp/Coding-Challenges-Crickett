CC = gcc
CFLAGS = -Wall -Wextra -std=c99

SRCS = frequency_counter.c huffman_tree.c
OBJS = $(SRCS:.c=.o)

TEST_SRCS = test_huffman_main.c
TEST_OBJS = $(TEST_SRCS:.c=.o)

MAIN_SRCS = main.c
MAIN_OBJS = $(MAIN_SRCS:.c=.o)

EXECUTABLE = huffman_compression
TEST_EXECUTABLE = test_huffman_tree

all: $(EXECUTABLE) $(TEST_EXECUTABLE)

$(EXECUTABLE): $(OBJS) $(MAIN_OBJS)
	$(CC) $(CFLAGS) $^ -o $@

$(TEST_EXECUTABLE): $(OBJS) $(TEST_OBJS)
	$(CC) $(CFLAGS) $^ -o $@

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

.PHONY: clean test run build

test: $(TEST_EXECUTABLE)
	./$(TEST_EXECUTABLE)

run: $(EXECUTABLE)
	./$(EXECUTABLE) input.txt

build: $(EXECUTABLE)

clean: 
	rm -f $(OBJS) $(TEST_OBJS) $(MAIN_OBJS) $(EXECUTABLE) $(TEST_EXECUTABLE)
