let count_bytes file_name = 
    let ic = open_in_bin file_name in
    let len = in_channel_length ic in
    close_in ic;
    len

let count_lines file_name = 
    let ic = open_in file_name in
    let count = ref 0 in
    try
        while true do 
            let _ = input_line ic in
            incr count
        done;
        !count
    with End_of_file -> 
        close_in ic;
        !count

let count_words file_name =
    let ic = open_in file_name in
    let count = ref 0 in
    try
        while true do 
            let line = input_line ic in
            let words = String.split_on_char ' ' line in
            count := !count + List.length words
        done;
        !count
    with End_of_file ->
        close_in ic;
        !count

let run_main () = 
    let bytes_flag = ref false in
    let lines_flag = ref false in
    let words_flag = ref false in
    let file_name = ref "" in
    let spec = [
        ("-c", Arg.Set bytes_flag, "Count bytes");
        ("-l", Arg.Set lines_flag, "Count lines");
        ("-w", Arg.Set words_flag, "Count words");
    ] in
    let usage = "Usage: wc [-c] [-l] <file>" in
    let anon_fun s = file_name := s in
    Arg.parse spec anon_fun usage;
    if !file_name = "" then
        (Printf.eprintf "%s\n" usage;
        exit 1)
    else
        try
            if !bytes_flag then
                let count = count_bytes !file_name in
                Printf.printf "%d %s\n" count !file_name
            else if !lines_flag then
                let count = count_lines !file_name in
                Printf.printf "%d %s\n" count !file_name
            else if !words_flag then
                let count = count_words !file_name in
                Printf.printf "%d %s\n" count !file_name
            else
                (Printf.eprintf "%s\n" usage;
                exit 1)
        with Sys_error _ ->
            Printf.eprintf "Error: Could not open file '%s'\n" !file_name;
            exit 1

let () = if Array.length Sys.argv > 1 then run_main ()
