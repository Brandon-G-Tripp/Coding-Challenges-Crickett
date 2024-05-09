defmodule ParserTest do 
  use ExUnit.Case
  doctest Parser

  test "parse empty object" do 
    valid_json = "{}"
    assert Parser.parse(valid_json) == {:ok, %{}}
  end

  test "parse object with single pair" do 
    json = ~s({"key": "value"})
    assert Parser.parse(json) == {:ok, %{"key" => "value"}}
  end

  test "parse object with multiple pairs" do 
    json = ~s({"key1": "value1", "key2": "value2"})
    assert Parser.parse(json) == {:ok, %{"key1" => "value1", "key2" => "value2"}}
  end

  test "parse invalid JSON" do 
    invalid_json = "invalid"
    assert Parser.parse(invalid_json) == {:error, "Invalid JSON object"}
  end

  test "parse incomplete JSON object" do 
    incomplete_json = "{"
    assert Parser.parse(incomplete_json) == {:error, "Invalid key-value pair"}
  end

  test "parse invalid object with trailing comma" do 
    json = ~s({"key": "value",})
    assert Parser.parse(json) == {:error, "Invalid trailing comma"}
  end

  test "parse object with different values types" do 
    json = ~s({"key1": true, "key2": false, "key3": null, "key4": "value", "key5": 101})
    assert Parser.parse(json) == {:ok, %{
      "key1" => true,
      "key2" => false, 
      "key3" => nil,
      "key4" => "value",
      "key5" => 101
    }}
  end

end
