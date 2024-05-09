defmodule CLI do 
  def main(args) do 
    case args do 
      [file_path] -> 
        parse_file(file_path)
      _ -> 
        IO.puts("Please provide a JSON file path as a command-line argument.")
        System.halt(1)
    end
  end

  defp parse_file(file_path) do 
    case File.read(file_path) do 
      {:ok, content} ->
        case Parser.parse(content) do 
          {:ok, _result} -> 
            IO.puts("Valid JSON")
            System.halt(0)
          
          {:error, reason} -> 
            IO.puts("Invalid JSON: #{reason}")
            System.halt(1)
        end

      {:error, reason} ->
        IO.puts("Error: #{reason}")
        System.halt(1)
    end
  end
end
