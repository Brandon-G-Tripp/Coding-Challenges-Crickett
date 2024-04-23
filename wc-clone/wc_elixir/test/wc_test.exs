defmodule WcTest do 
  use ExUnit.Case
  doctest Wc

  test "count_bytes/1 counts the number of bytes in a file" do 
    content = "Sample content"
    file_name = "test.txt"
    File.write!(file_name, content)
    assert Wc.count_bytes(file_name) == byte_size(content)
    File.rm!(file_name)
  end

  test "count_bytes/1 returns an error for a non-existent file" do
    file_name = "nonexistent.txt"
    assert_raise File.Error, fn -> 
      Wc.count_bytes(file_name)
    end
  end

end

