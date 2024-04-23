defmodule Wc do 
  def main(args) do 
    {opts, [file_name], _} = OptionParser.parse(args, switches: [bytes: :boolean], aliases: [c: :bytes])

    if opts[:bytes] do 
      case File.read(file_name) do 
        {:ok, content} -> 
          IO.puts("#{byte_size(content)} #{file_name}")
        {:error, _} ->
          IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
          System.halt(1)
      end
    else
      IO.puts(:stderr, "Error: Missing -c flag")
      System.halt(1)
    end
  end

  def count_bytes(file_name) do 
    File.read!(file_name)
    |> byte_size()
  end
end
