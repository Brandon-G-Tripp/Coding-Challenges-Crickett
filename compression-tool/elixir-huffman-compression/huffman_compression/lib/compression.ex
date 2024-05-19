defmodule Compression do 
  def count_frequencies(text) do 
    text 
    |> String.graphemes()
    |> Enum.reduce(%{}, fn char, acc ->
      Map.update(acc, char, 1, &(&1 +1))
    end)
  end
end
