defmodule ElixirSort.Sorter do
  def sort_file(filename, opts \\ []) do 
    unique = Keyword.get(opts, :unique, false)

    filename
    |> File.stream!()
    |> Enum.map(&String.trim/1)
    |> then(fn lines ->
      if unique do  
        Enum.uniq(lines)
      else
        lines
      end
    end)
    |> Enum.sort()
  end
end
