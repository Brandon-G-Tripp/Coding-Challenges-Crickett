let count_bytes file_name = 
    let ic = open_in_bin file_name in
    let len = in_channel_length ic in
    close_in ic;
    len


let main () = 
    let bytes_flag = ref false in
    let spec = [("-c", Arg.Set bytes_flag, "Count bytes")] in
    let usage = "Usage: wc [-c] <file>" in
    Arg.parse spec (fun _ -> ()) usage;
    let file_name = Sys.argv.(Array.length Sys.argv - 1) in
    if !bytes_flag then
        try 
            let count = count_bytes file_name in
            Printf.printf "%d %s\n" count file_name
        with Sys_error _ -> 
            Printf.eprintf "Error: Could not open file '%s'\n" file_name;
            exit 1
    else
        (Printf.eprintf "Error: Missing -c flag \n";
        exit 1)

let () = main ()
