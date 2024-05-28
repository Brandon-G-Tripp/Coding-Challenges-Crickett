defmodule CutTool do 
  def cut_second_field(file_path) do 
    file_path
    |> File.stream!()
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.split(&1, "\t"))
    |> Stream.map(fn fields -> Enum.at(fields, 1) end)
  end
end
