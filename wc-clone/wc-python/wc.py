import argparse 
import sys


def default_count(file_path):
    try:
        line_count = count_lines(file_path)
        word_count = count_words(file_path)
        char_count = count_chars(file_path)
        return line_count, word_count, char_count
    except FileNotFoundError:
        return None, None, None


def count_chars(file_path):
    try:
        with open(file_path, "r", encoding="utf-8") as file:
            content = file.read()
            return len(content)
    except FileNotFoundError:
        return None

def count_bytes(file_path):
    try:
        with open(file_path, "rb") as file:
            return len(file.read())
    except FileNotFoundError:
        return None

def count_lines(file_path):
    try:
        with open(file_path, "r") as file:
            return len(file.readlines())
    except FileNotFoundError:
        return None

def count_words(file_path):
    try:
        with open(file_path, "r") as file:
            content = file.read()
            words = content.split()
            return len(words)
    except FileNotFoundError:
        return None

def main():
    parser = argparse.ArgumentParser(description="Word Count")
    parser.add_argument("-c", "--bytes", action="store_true", help="Count bytes")
    parser.add_argument("-l", "--lines", action="store_true", help="Count lines")
    parser.add_argument("-w", "--words", action="store_true", help="Count words")
    parser.add_argument("-m", "--chars", action="store_true", help="Count characters")
    parser.add_argument("file", help="Input file")
    args = parser.parse_args()

    if not any([args.bytes, args.lines, args.words, args.chars]):
        line_count, word_count, char_count = default_count(args.file)
        if line_count is None:
            print(f"Error: could not open file '{args.file}'")
            sys.exit(1)
        print(f"{line_count} {word_count} {char_count} {args.file}")
    else:
        if args.chars:
            count = count_chars(args.file)
            if count is None:
                print(f"Error: could not open file '{args.file}'")
                sys.exit(1)
            print(f"{count} {args.file}")
        elif args.bytes:
            count = count_bytes(args.file)
            if count is None:
                print(f"Error: could not open file '{args.file}'")
                sys.exit(1)
            print(f"{count} {args.file}")
        elif args.lines:
            count = count_lines(args.file)
            if count is None: 
                print(f"Error: could not open file '{args.file}'")
                sys.exit(1)
            print(f"{count} {args.file}")
        elif args.words:
            count = count_words(args.file)
            if count is None: 
                print(f"Error: could not open file '{args.file}'")
                sys.exit(1)
            print(f"{count} {args.file}")
        else:
            print("Error: Missing -c or -l, or -w flag")
            sys.exit(1)


if __name__ == "__main__":
    main()

