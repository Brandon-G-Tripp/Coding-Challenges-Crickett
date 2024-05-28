defmodule CutTool.CLI do 
  def main(args) do 
    case args do 
      ["-f2", file_path] -> 
        file_path 
        |> CutTool.cut_second_field()
        |> Enum.each(&IO.puts/1)

      _ -> 
        IO.puts("Usage: cut_tool -f <file_path>")
    end
  end
end
