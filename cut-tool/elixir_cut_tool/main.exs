defmodule Main do 
  def main(args) do
    case args do 
      ["-f2", file_path] -> 
        file_path
        |> Cut.cut_second_field()
        |> Enum.each(&IO.puts/1)

      _ -> 
        IO.puts("Usage: elixir main.exs -f2 <file_path>")
    end
  end
end


Main.main(System.argv())
