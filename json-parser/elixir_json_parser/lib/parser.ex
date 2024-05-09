defmodule Parser do
  def parse(json) do
    json = String.trim(json)

    case json do
      "{}" ->
        {:ok, %{}}

      "{" <> rest ->
        parse_object(rest, %{})

      _ ->
        {:error, "Invalid JSON object"}
    end
  end

  defp parse_object("}", acc), do: {:ok, acc}

  defp parse_object(json, acc) do
    case parse_key_value(json) do
      {:ok, key, value, rest} ->
        acc = Map.put(acc, key, value)
        case String.trim_leading(rest) do
          "," <> "}" <> _ ->
            {:error, "Invalid trailing comma"}

          "," <> rest ->
            parse_object(String.trim_leading(rest), acc)

          "}" ->
            {:ok, acc}

          _ ->
            {:error, "Invalid JSON object"}
        end

      {:error, reason} ->
        {:error, reason}
    end
  end

  defp parse_key_value(json) do
    with {:ok, key, ":" <> rest} <- parse_string(json),
         {:ok, value, rest} <- parse_value(String.trim_leading(rest)) do
      {:ok, key, value, String.trim_leading(rest)}
    else
      _ -> {:error, "Invalid key-value pair"}
    end
  end

  defp parse_string(""), do: {:error, "Invalid string"}
  defp parse_string(json) do
    case String.trim(json) do
      "\"" <> rest ->
        case String.split(rest, "\"", parts: 2) do
          [value, rest] ->
            rest = String.trim_leading(rest)
            case rest do
              ":" <> rest -> {:ok, value, String.trim_leading(rest)}
              _ -> {:error, "Invalid string"}
            end
          _ ->
            {:error, "Invalid string"}
        end
      _ ->
        {:error, "Invalid string"}
    end
  end

  defp parse_value("true" <> rest), do: {:ok, true, String.trim_leading(rest)}
  defp parse_value("false" <> rest), do: {:ok, false, String.trim_leading(rest)}
  defp parse_value("null" <> rest), do: {:ok, nil, String.trim_leading(rest)}
  defp parse_value("\"" <> _ = json), do: parse_string(json)
  defp parse_value(<<char, _::binary>> =json) when char in ?0..?9 or char == ?-, do: parse_number(json)
  defp parse_value(_), do: {:error, "Invalid JSON value"}

  defp parse_number(json) do
    case Float.parse(json) do
      {value, rest} -> {:ok, value, rest}
      :error -> {:error, "Invalid number"}
    end
  end
end
