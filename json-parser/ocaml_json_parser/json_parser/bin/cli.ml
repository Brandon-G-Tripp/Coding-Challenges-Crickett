open Json_parser.Parser

let () = 
    if Array.length Sys.argv < 2 then (
        Printf.eprintf "Please provide a JSON file path as a command line argument.\n";
        exit 1
    );

let file_path = Sys.argv.(1) in

try 
    let ch = open_in file_path in
    let file_content = really_input_string ch (in_channel_length ch) in
    close_in ch;

    let result = parse_json file_content in
    if result then (
        Printf.printf "Valid JSON\n";
        exit 0 
    ) else (
        Printf.printf "Invalid JSON\n";
        exit 1
    )
with Sys_error _ -> 
    Printf.eprintf "Error: JSON file not found.\n";
    exit 1
