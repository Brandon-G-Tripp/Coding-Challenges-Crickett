import sys 

def cut_second_field(file_path):
    with open(file_path, "r") as file: 
        for line in file: 
            fields = line.strip().split("\t")
            if len(fields) >= 2:
                print(fields[1])
