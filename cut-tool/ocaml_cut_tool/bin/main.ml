open Cut

let () = 
    if Array.length Sys.argv <> 3 || Sys.argv.(1) <> "-f2" then 
        Printf.eprintf "Usage: %s -f2 <file_path>\n" Sys.argv.(0)
    else
        let file_path = Sys.argv.(2) in
        let output = cut_second_field file_path in 
        List.iter print_endline output
