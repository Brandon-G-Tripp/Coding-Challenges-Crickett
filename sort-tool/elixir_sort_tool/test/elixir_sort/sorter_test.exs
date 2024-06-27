defmodule ElixirSort.SorterTest do 
  use ExUnit.Case
  alias ElixirSort.Sorter

  test "sort_file sorts lines in a file lexicographically" do 
    input = """
    banana
    apple
    cherry
    """
    filename = "test_input.txt"
    File.write!(filename, input)

    result = Sorter.sort_file(filename)

    expected = ["apple", "banana", "cherry"]
    assert result == expected

    File.rm!(filename)
  end
end
