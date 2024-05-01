import argparse 
import io
import sys


def count_bytes(file):
    if file is sys.stdin:
        content = file.read().encode('utf-8')
        return len(content)
    else:
        try:
            with open(file, "rb") as f:
                content = f.read()
                return len(content)
        except FileNotFoundError:
            return None


def count_chars(file):
    if file is sys.stdin:
        content = file.read()
        return len(content)
    else:
        try:
            with open(file, "r", encoding="utf-8") as f:
                content = f.read()
                return len(content)
        except FileNotFoundError:
            return None

def count_lines(file):
    if file is sys.stdin:
        return sum(1 for _ in file)
    else:
        try:
            with open(file, "r") as f:
                return sum(1 for _ in f)
        except FileNotFoundError:
            return None

def count_words(file):
    if file is sys.stdin:
        content = file.read()
        words = content.split()
        return len(words)
    else:
        try:
            with open(file, "r") as f:
                content = f.read()
                words = content.split()
                return len(words)
        except FileNotFoundError:
            return None

def default_count(file):
    try:
        if file is sys.stdin:
            content = file.read()
            line_count = content.count('\n')
            if content and not content.endswith('\n'):
                line_count += 1
            word_count = len(content.split())
            char_count = len(content)
        else:
            with open(file, "r") as f:
                content = f.read()
                line_count = content.count('\n')
                if content and not content.endswith('\n'):
                    line_count += 1
                word_count = len(content.split())
                char_count = len(content)
        return line_count, word_count, char_count
    except FileNotFoundError:
        return None, None, None


def main():
    parser = argparse.ArgumentParser(description="Word Count")
    parser.add_argument("-c", "--bytes", action="store_true", help="Count bytes")
    parser.add_argument("-l", "--lines", action="store_true", help="Count lines")
    parser.add_argument("-w", "--words", action="store_true", help="Count words")
    parser.add_argument("-m", "--chars", action="store_true", help="Count characters")
    parser.add_argument("file", nargs="?", default=None, help="Input file")
    args = parser.parse_args()

    if args.file == "-" or args.file is None:
        file = sys.stdin
    else:
        file = args.file

    if not any([args.bytes, args.lines, args.words, args.chars]):
        line_count, word_count, char_count = default_count(file)
        if line_count is None:
            print(f"Error: could not open file '{args.file}'")
            sys.exit(1)
        print(f"{line_count} {word_count} {char_count} {args.file}")
    else:
        if args.chars:
            count = count_chars(file)
            print(f"{count}")
        elif args.bytes:
            count = count_bytes(file)
            print(f"{count}")
        elif args.lines:
            count = count_lines(file)
            print(f"{count}")
        elif args.words:
            count = count_words(file)
            print(f"{count}")
        else:
            print("Error: Missing -c or -l, or -w flag")
            sys.exit(1)


if __name__ == "__main__":
    main()

