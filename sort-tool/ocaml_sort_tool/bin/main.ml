open Ocaml_sort_tool.Sort

let () =
    match Array.to_list Sys.argv with
    | [_; "-u"; filename] ->
        let sorted_lines = sort_file ~unique:true filename in
        List.iter print_endline sorted_lines
    | [_; filename] -> 
        let sorted_lines = sort_file filename in
        List.iter print_endline sorted_lines
    | _ -> 
            Printf.printf "Usage: %s <filename>\n" Sys.argv.(0)
