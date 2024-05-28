defmodule CutTest do 
  use ExUnit.Case

  test "cut_second_field/1 prints the second field of each line" do 
    # Prepare the input file
    input_data = "a\tf1\t1\nb\t1\t2\nc\t6\t3\nd\t11\t4\ne\t16\t5\nf\t21\t6\n"
    File.write!("sample.tsv", input_data)

    # Run the cut command and capture the output 
    output =
      "sample.tsv"
      |> CutTool.cut_second_field()
      |> Enum.join("\n")

    # Check the output 
    expected_output = "f1\n1\n6\n11\n16\n21"
    assert output == expected_output

    # Clean up the temporary file
    File.rm!("sample.tsv")
  end
end
