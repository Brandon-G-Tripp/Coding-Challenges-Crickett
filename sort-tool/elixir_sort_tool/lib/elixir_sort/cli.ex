defmodule ElixirSort.CLI do 
  def main(args) do 
    case args do 
      ["-u", filename] -> 
        sorted_lines = ElixirSort.Sorter.sort_file(filename, unique: true)
        IO.write(Enum.join(sorted_lines, "\n") <> "\n")
      [filename] -> 
        sorted_lines = ElixirSort.Sorter.sort_file(filename)
        IO.write(Enum.join(sorted_lines, "\n") <> "\n")
      _ -> 
        IO.puts("Usage: elixir_sort [-u] <filename>")
    end
  end
end
