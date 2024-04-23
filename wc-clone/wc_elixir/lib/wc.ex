defmodule Wc do 
  def count_bytes(file_name) do 
    File.read!(file_name)
    |> byte_size()
  end
end
