defmodule Wc do
  def main(args) do
    {opts, file_names, _} =
      OptionParser.parse(
        args,
        switches: [bytes: :boolean, lines: :boolean, words: :boolean, chars: :boolean],
        aliases: [c: :bytes, l: :lines, w: :words, m: :chars]
      )

    case {opts, file_names} do
      {[], []} ->
        IO.puts(:stderr, "Usage: wc [-c] [-l] [-w] [-m] <file>")
        System.halt(1)

      {opts, []} ->
        IO.puts(:stderr, "Usage: wc [-c] [-l] [-w] [-m]  <file>")
        System.halt(1)

      {[], [file_name]} -> 
        case default_count(file_name) do
          {:ok, line_count, word_count, byte_count} ->
            IO.puts("#{line_count} #{word_count} #{byte_count} #{file_name}")
          {:error, _} ->
            IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
            System.halt(1)
        end

      {opts, [file_name]} ->
        cond do
          opts[:bytes] ->
            case Wc.count_bytes(file_name) do
              {:ok, count} ->
                IO.puts("#{count} #{file_name}")
              {:error, _} ->
                IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
                System.halt(1)
            end
          opts[:lines] ->
            case Wc.count_lines(file_name) do
              {:ok, count} ->
                IO.puts("#{count} #{file_name}")
              {:error, _} ->
                IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
                System.halt(1)
            end
          opts[:words] ->
            case Wc.count_words(file_name) do
              {:ok, count} ->
                IO.puts("#{count} #{file_name}")
              {:error, _} ->
                IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
                System.halt(1)
            end
          opts[:chars] ->
            case Wc.count_chars(file_name) do
              {:ok, count} ->
                IO.puts("#{count} #{file_name}")
              {:error, _} ->
                IO.puts(:stderr, "Error: Could not open file '#{file_name}'")
                System.halt(1)
            end
          true ->
            IO.puts(:stderr, "Usage: wc [-c] [-l] <file>")
            System.halt(1)
        end

      {_, _} ->
        IO.puts(:stderr, "Usage: wc [-c] [-l] <file>")
        System.halt(1)
    end
  end

  defp default_count(file_name) do
    with {:ok, line_count} <- count_lines(file_name),
      {:ok, word_count} <- count_words(file_name),
      {:ok, byte_count} <- count_bytes(file_name) do
      {:ok, line_count, word_count, byte_count}
    else
    {:error,_} -> {:error, :file_not_found}
    end
  end

  def count_bytes(file_name) do
    try do
      content = File.read!(file_name)
      {:ok, byte_size(content)}
    rescue
      File.Error -> {:error, :file_not_found}
    end
  end

  def count_lines(file_name) do
    try do
      count = File.stream!(file_name)
              |> Enum.count()
      {:ok, count}
    rescue
      File.Error -> {:error, :file_not_found}
    end
  end

  def count_words(file_name) do 
    try do 
      count = File.stream!(file_name)
        |> Stream.map(&String.split/1)
        |> Stream.map(&length/1)
        |> Enum.sum()
      {:ok, count}
    rescue
      File.Error -> {:error, :file_not_found}
    end
  end

  def count_chars(file_name) do 
    try do 
      content = File.read!(file_name)
      {:ok, String.length(content)}
    rescue
    File.Error -> {:error, :file_not_found}
    end
  end
end
