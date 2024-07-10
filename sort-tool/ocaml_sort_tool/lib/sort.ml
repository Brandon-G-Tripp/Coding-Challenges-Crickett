let sort_file ?(unique=false) filename = 
    let chan = open_in filename in
    let lines =  ref [] in
    try
        while true do 
            lines := (input_line chan) :: !lines
        done;
        []
    with End_of_file -> 
        close_in chan;
        let sorted = List.rev !lines |> List.sort String.compare in
        if unique then
            List.sort_uniq String.compare sorted
        else
            sorted
