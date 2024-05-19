open Compression

let () = 
    if Array.length Sys.argv <> 2 then (
        Printf.printf "Usage: %s <file_path>\n" Sys.argv.(0);
        exit 1
    );
    let file_path = Sys.argv.(1) in
    let text = In_channel.with_open_text file_path In_channel.input_all in
    let frequencies = count_frequencies text in 
    Printf.printf "Character Frequencies:\n";
    Hashtbl.iter (fun char count -> 
        Printf.printf "%c: %d\n" char count
    ) frequencies
