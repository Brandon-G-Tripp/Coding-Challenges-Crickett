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
    case String.trim(input) do 
      "}" <> rest -> {:ok, %{}, String.trim(rest)}
      _ -> {:error, "Invalid JSON object"}
    end
  end
end
