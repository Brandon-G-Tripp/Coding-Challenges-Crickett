defmodule ElixirSort.Sorter do
  def sort_file(filename) do 
    filename
    |> File.stream!()
    |> Enum.map(&String.trim/1)
    |> Enum.sort()
  end
end
