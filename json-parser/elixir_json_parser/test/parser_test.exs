defmodule ParserTest do 
  use ExUnit.Case
  doctest Parser

  test "parse empty object" do 
    valid_json = "{}"
    assert Parser.parse_json(valid_json) == true
  end

  test "parse object with single pair" do 
    json = ~s({"key": "value"})
    assert Parser.parse_json(json) == true
  end

  test "parse object with multiple pairs" do 
    json = ~s({"key1": "value1", "key2": "value2"})
    assert Parser.parse_json(json) == true
  end

  test "parse invalid JSON" do 
    invalid_json = "invalid"
    assert Parser.parse_json(invalid_json) == false
  end

  test "parse incomplete JSON object" do 
    incomplete_json = "{"
    assert Parser.parse_json(incomplete_json) == false
  end

  test "parse invalid object with trailing comma" do 
    json = ~s({"key": "value",})
    assert Parser.parse_json(json) == false
  end

end
