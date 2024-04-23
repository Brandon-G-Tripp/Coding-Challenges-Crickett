import argparse 
import sys

def count_bytes(file_path):
    try:
        with open(file_path, "rb") as file:
            return len(file.read())
    except FileNotFoundError:
        return None


def main():
    parser = argparse.ArgumentParser(description="Word Count")
    parser.add_argument("-c", "--bytes", action="store_true", help="Count bytes")
    parser.add_argument("file", help="Input file")
    args = parser.parse_args()

    if args.bytes:
        count = count_bytes(args.file)
        if count is None:
            print(f"Error: could not open file '{args.file}'")
            sys.exit(1)
        print(f"{count} {args.file}")
    else:
        print("Error: Missing -c flag")
        sys.exit(1)


if __name__ == "__main__":
    main()

