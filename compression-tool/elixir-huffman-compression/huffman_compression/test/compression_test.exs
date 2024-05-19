defmodule CompressionTest do 
  use ExUnit.Case 
  doctest Compression

  test "count_frequencies/1 counts the frequenciesof characters correctly" do 
    text = "aabbccabc"
    expected_frequencies = %{"a" => 3, "b" => 3, "c" => 3}
    assert Compression.count_frequencies(text) == expected_frequencies
  end
end
