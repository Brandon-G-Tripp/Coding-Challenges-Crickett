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
        process_input(:stdio, opts)

      {opts, []} ->
        process_input(:stdio, opts)

      {[], [file_name]} -> 
        case default_count(file_name) do
          {:ok, [line_count, word_count, byte_count]} ->
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

  defp default_count(input) do
    case input do 
      :stdio -> 
        content = IO.read(:stdio, :all)
        line_count = content |> String.split("\n", trim: true) |> length()
        word_count = content |> String.split() |> length()
        byte_count = byte_size(content)
        {:ok, [line_count, word_count, byte_count]}
      file_name -> 
        case File.read(file_name) do 
          {:ok, content} -> 
            line_count = content |> String.split("\n", trim: true) |> length()
            word_count = content |> String.split() |> length()
            byte_count = byte_size(content)
            {:ok, [line_count, word_count, byte_count]}
          {:error, _ } ->
            {:error, :file_not_found}
        end
    end
  end

  def count_bytes(file_name) do
    case File.read(file_name) do
      {:ok, content} -> {:ok, byte_size(content)}
      {:error, _} -> {:error, :file_not_found}
    end
  end

  def count_lines(file_name) do
    case File.read(file_name) do
      {:ok, content} ->
        {:ok, content |> String.trim() |> String.split("\n") |> length()}
      {:error, _} ->
        {:error, :file_not_found}
    end
  end

  def count_words(file_name) do
    case File.read(file_name) do
      {:ok, content} ->
        {:ok, content |> String.split() |> length()}
      {:error, _} ->
        {:error, :file_not_found}
    end
  end

  def count_chars(file_name) do
    case File.read(file_name) do
      {:ok, content} -> {:ok, String.length(content)}
      {:error, _} -> {:error, :file_not_found}
    end
  end

  defp process_input(:stdio, opts) do
    input = IO.read(:stdio, :all)

    result =
      cond do
        opts[:bytes] -> {:ok, byte_size(input)}
        opts[:lines] -> {:ok, input |> String.split("\n") |> length()}
        opts[:words] -> {:ok, input |> String.split() |> length()}
        opts[:chars] -> {:ok, String.length(input)}
        true ->
          line_count = input |> String.trim() |> String.split("\n") |> length()
          word_count = input |> String.trim() |> String.split() |> length()
          byte_count = byte_size(input)
          {:ok, [line_count, word_count, byte_count]}
      end

    case result do
      {:ok, [line_count, word_count, byte_count]} ->
        IO.puts("#{line_count} #{word_count} #{byte_count}")
      {:ok, counts} ->
        case counts do
          [line_count, word_count, byte_count] ->
            IO.puts("#{line_count} #{word_count} #{byte_count}")
          count ->
            IO.puts(count)
        end
      {:error, _} ->
        System.halt(1)
    end
  end

end
