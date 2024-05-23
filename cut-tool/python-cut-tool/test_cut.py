import sys 

def test_cut_second_field():
    # Prepare the input file
    input_data = "a\tf1\t1\nb\t1\t2\nc\t6\t3\nd\t11\t4\ne\t16\t5\nf\t21\t6\n"
    with open("sample.tsv", "w") as file:
        file.write(input_data)

    # Run the cut command and capture the output 
    import subprocess
    output = subprocess.check_output([sys.executable, "main.py", "-f2", "sample.tsv"]).decode().strip()

    # Check the output
    expected_output = "f1\n1\n6\n11\n16\n21"
    assert output == expected_output

    # Clean up the temporary file
    import os 
    os.remove("sample.tsv")
