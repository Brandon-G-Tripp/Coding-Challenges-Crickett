import argparse 
import sys

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

def main():
    parser = argparse.ArgumentParser(description="Word Count")
    parser.add_argument("-c", "--bytes", action="store_true", help="Count bytes")
    parser.add_argument("-l", "--lines", action="store_true", help="Count lines")
    parser.add_argument("file", help="Input file")
    args = parser.parse_args()

    if args.bytes:
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
    else:
        print("Error: Missing -c or -l flag")
        sys.exit(1)


if __name__ == "__main__":
    main()

