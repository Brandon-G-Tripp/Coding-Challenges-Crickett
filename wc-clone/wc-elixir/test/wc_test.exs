defmodule WcTest do 
  use ExUnit.Case
  import ExUnit.CaptureIO
  doctest Wc

  test "main/1 counts lines, words, and bytes when no option is provided" do 
    content = "Line 1\nLine 2\nLine 3\n"
    file_name = "test.txt"
    File.write!(file_name, content)
    expected_output = "3 6 21 test.txt\n"
    assert capture_io(fn -> Wc.main([file_name]) end) == expected_output
    File.rm!(file_name)
  end

  test "count_bytes/1 counts the number of bytes in a file" do 
    content = "Sample content"
    file_name = "test.txt"
    File.write!(file_name, content)
    assert Wc.count_bytes(file_name) == {:ok, byte_size(content)}
    File.rm!(file_name)
  end

  test "count_bytes/1 returns an error for a non-existent file" do
    file_name = "nonexistent.txt"
    assert Wc.count_bytes(file_name) == {:error, :file_not_found}
  end

  test "count_lines/1 counts the number of lines in a file" do 
    content = "Line 1\nLine 2\nLine 3\n"
    file_name = "test.txt"
    File.write!(file_name, content)
    assert Wc.count_lines(file_name) == {:ok, 3}
    File.rm!(file_name)
  end

  test "count_lines/1 returns an error for a non-existent file" do 
    file_name = "nonexistent.txt"
    assert Wc.count_lines(file_name) == {:error, :file_not_found}
  end

  test "count_words/1 counts the number of words in a file" do 
    content = "This is a sample file\nwith multiple words\non each line\n"
    file_name = "test.txt"
    File.write!(file_name, content)
    assert Wc.count_words(file_name) == {:ok, 11}
    File.rm!(file_name)
  end

  test "count_words/1 returns an error for a non-existent file" do 
    file_name = "nonexistent.txt"
    assert Wc.count_words(file_name) == {:error, :file_not_found}
  end

  test "count_chars/1 counts the number of characters in a file" do 
    content = "Sample content with 🚀 emoji"
    file_name = "test.txt"
    File.write!(file_name, content)
    assert Wc.count_chars(file_name) == {:ok, String.length(content)}
    File.rm!(file_name)
  end

  test "count_chars/1 returns an error for a non-existent file" do 
    file_name = "nonexistent.txt"
    assert Wc.count_chars(file_name) == {:error, :file_not_found}
  end

end

