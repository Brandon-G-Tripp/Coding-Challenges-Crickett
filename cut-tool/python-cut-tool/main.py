import sys 
from cut import cut_second_field

def main():
    if len(sys.argv) == 3 and sys.argv[1] == "-f2": 
        file_path = sys.argv[2]
        cut_second_field(file_path)
    else:
        print("Usage: python main.py -f2 <file_path>")

if __name__ == "__main__":
    main()
