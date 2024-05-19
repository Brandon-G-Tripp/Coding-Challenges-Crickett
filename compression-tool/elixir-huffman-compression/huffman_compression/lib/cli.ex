defmodule Compression.CLI do 
  def main(args) do 
    case args do 
      [file_path] -> 
        file_path
        |> File.read!()
        |> Compression.count_frequencies()
        |> IO.inspect(lable: "Character Frequencies")

      _ -> 
        IO.puts("Usage: mix run lib/cli.ex <file_path>")
    end
  end
end
