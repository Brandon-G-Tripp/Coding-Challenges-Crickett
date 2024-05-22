CC = gcc
CFLAGS = -Wall -Wextra -Werror

.PHONY: build clean test

build: cut

cut: main.o cut.o
	$(CC) $(CFLAGS) -o cut main.o cut.o

main.o: main.c cut.h
	$(CC) $(CFLAGS) -c main.c

cut.o: cut.c cut.h
	$(CC) $(CFLAGS) -c cut.c

clean:
	rm -f cut *.o

test: test_main.o test_cut.o cut.o
	$(CC) $(CFLAGS) -o test_cut test_main.o test_cut.o cut.o
	./test_cut
