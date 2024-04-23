defmodule WcElixirTest do
  use ExUnit.Case
  doctest WcElixir

  test "greets the world" do
    assert WcElixir.hello() == :world
  end
end
