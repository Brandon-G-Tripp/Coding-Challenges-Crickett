defmodule ParserTest do 
  use ExUnit.Case
  doctest Parser

  test "parse valid JSON" do 
    valid_json = "{}"
    assert Parser.parse_json(valid_json) == true
  end

  test "parse invalid JSON" do 
    invalid_json = "invalid"
    assert Parser.parse_json(invalid_json) == false
  end

  test "parse incomplete JSON object" do 
    incomplete_json = "{"
    assert Parser.parse_json(incomplete_json) == false
  end

end
