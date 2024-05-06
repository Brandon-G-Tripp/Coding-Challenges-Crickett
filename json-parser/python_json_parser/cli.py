import sys 
from src.parser import parse_json

if len(sys.argv) < 2:
    print('Please provide a JSON file path as a command line argument.')
    sys.exit(1)

file_path = sys.argv[1]

try: 
    with open(file_path, 'r') as file:
        file_content = file.read()
        result = parse_json(file_content)

        if result: 
            print('Valid JSON')
            sys.exit(0)
        else:
            print('Invalid JSON')
            sys.exit(1)
except FileNotFoundError:
    print('Error: JSON file not found.')
    sys.exit(1)
