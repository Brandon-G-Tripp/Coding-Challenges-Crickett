defmodule Parser do 
  def parse_json(input) do 
    case parse_value(String.trim(input)) do
      {:ok, _value, ""} -> true
      _ -> false
    end
  end


  defp parse_value("{" <> rest), do: parse_object(rest)
  defp parse_value(_), do: {:error, "Invalid JSON value"}

  defp parse_object(input) do
    case parse_pairs(String.trim(input)) do 
      {:ok, pairs, "}" <> rest} -> {:ok, Map.new(pairs), String.trim(rest)}
      _ -> {:error, "Invalid JSON object"}
    end
  end

  defp parse_pairs(input) do 
    case String.trim(input) do 
      "}" <> rest -> 
        {:ok, [], "}" <> String.trim(rest)}
      _-> 
        case parse_pair(input) do
          {:ok, pair, "," <> rest} ->
            case String.trim(rest) do 
              "}" <> _ -> {:error, "Invalid JSON object"}
              _ -> 
                case parse_pairs(String.trim(rest)) do
                  {:ok, pairs, rest} -> {:ok, [pair | pairs], rest}
                  _ -> {:error, "Invalid JSON object"}
                end
              end
          {:ok, pair, rest} -> 
            case String.trim(rest) do
              "}" <> _ -> {:ok, [pair], rest}
              _ -> {:error, "Invalid JSON object"}
            end
          _ -> 
            case String.trim(input) do
              "}" <> rest -> {:ok, [], rest}
              _ -> {:error, "Invalid JSON object"}
            end
        end
    end
  end

  defp parse_pair(input) do
    with {:ok, key, ":" <> rest} <- parse_string(input),
        {:ok, value, rest} <- parse_string(String.trim(rest)) do
      {:ok, {key, value}, rest} 
    else
      _ -> {:error, "Invalid JSON pair"}
    end
  end

  defp parse_string("\"" <> rest) do 
    case Regex.run(~r/^((?:[^"\\]|\\.)*)"/, rest) do
      [match, string] -> {:ok, string, String.slice(rest, String.length(match)..-1)}
      nil -> {:error, "Invalid JSON string"}
    end
  end
  defp parse_string(_), do: {:error, "Invalid JSON string"}
end
