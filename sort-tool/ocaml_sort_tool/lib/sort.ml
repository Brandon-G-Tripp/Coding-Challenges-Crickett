let sort_file filename = 
    let chan = open_in filename in
    let lines =  ref [] in
    try
        while true do 
            lines := input_line chan :: !lines
        done;
        []
    with End_of_file -> 
        close_in chan;
        List.rev !lines |> List.sort String.compare
