defmodule ElixirSort.CLI do 
  def main(args) do 
    case args do 
      [filename] -> 
        sorted_lines = ElixirSort.Sorter.sort_file(filename)
        IO.write(Enum.join(sorted_lines, "\n") <> "\n")
      _ -> 
        IO.puts("Usage: elixir_sort <filename>")
    end
  end
end
