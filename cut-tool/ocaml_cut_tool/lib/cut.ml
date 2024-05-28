let cut_second_field file_path = 
    let ic = open_in file_path in
    let rec loop acc = 
        try
            let line = input_line ic in
            let fields = String.split_on_char '\t' line in
            match fields with 
            | _ :: field2 :: _ -> loop (field2 :: acc)
            | _ -> loop acc
        with End_of_file -> 
            close_in ic;
            List.rev acc
    in
    loop []
