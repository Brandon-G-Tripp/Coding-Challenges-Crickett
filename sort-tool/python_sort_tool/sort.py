import sys
import os

def sort_file(filename):
    try:
        with open(filename, 'r') as file:
            lines = file.readlines()
        if not lines:
            return ''
        # Sort lines case insensitively.
        sorted_lines = sorted(line.strip().upper() for line in lines)
        return '\n'.join(sorted_lines) + '\n'
    except IOError as e:
        print(f"Error reading file: {e}", file=sys.stderr)
        sys.exit(1)


def main():
    if len(sys.argv) != 2:
        print("Usage: python sort.py <filename>", file=sys.stderr)
        sys.exit(1)
    result = sort_file(sys.argv[1])
    try:
        print(result, end='')
    except BrokenPipeError:
        # Python flushes standard streams on exit; redirect remaining output
        # to devnull to avoid another BrokenPipeError at shutdown
        devnull = os.open(os.devnull, os.O_WRONLY)
        os.dup2(devnull, sys.stdout.fileno())
        sys.exit(0)

if __name__ == "__main__":
    main()
