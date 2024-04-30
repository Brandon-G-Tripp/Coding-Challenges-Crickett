let count_chars file_name = 
    let ic = open_in file_name in
    let count = ref 0 in 
    let decoder = Uutf.decoder ~encoding: `UTF_8 (`Channel ic) in
    let rec loop () = 
        match Uutf.decode decoder with
        | `Await -> close_in ic; !count
        | `End -> close_in ic; !count
        | `Malformed _ -> close_in ic; raise (Failure "Malformed UTF-8 data")
        | `Uchar _ -> 
                incr count;
                loop () 
    in
    loop ()

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

let run_main args = 
    let bytes_flag = ref false in
    let lines_flag = ref false in
    let words_flag = ref false in
    let chars_flag = ref false in
    let file_name = ref "" in
    let spec = [
        ("-c", Arg.Set bytes_flag, "Count bytes");
        ("-l", Arg.Set lines_flag, "Count lines");
        ("-w", Arg.Set words_flag, "Count words");
        ("-m", Arg.Set chars_flag, "Count characters");
    ] in
    let usage = "Usage: wc [-c] [-l] [-w] [-m] <file>" in
    let anon_fun s = file_name := s in
    Arg.parse_argv args spec anon_fun usage;
    if !file_name = "" then
        (Printf.eprintf "%s\n" usage;
        exit 1)
    else
        try
            if not !bytes_flag && not !lines_flag && not !words_flag && not !chars_flag then
                let line_count = count_lines !file_name in
                let word_count = count_words !file_name in
                let byte_count = count_bytes !file_name in
                Printf.printf "%d %d %d %s\n" line_count word_count byte_count !file_name
            else if !chars_flag then
                let count = count_chars !file_name in
                Printf.printf "%d %s\n" count !file_name
            else if !bytes_flag then
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

let () = if Array.length Sys.argv > 1 then run_main Sys.argv else ()
