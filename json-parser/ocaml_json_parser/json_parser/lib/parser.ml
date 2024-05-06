let parse_json input_string = 
    let index = ref 0 in
    let len = String.length input_string in 

    let consume_whitespace () = 
        while !index < len && String.contains " \t\n\r" input_string.[!index] do 
            incr index
        done
    in

    let consume_char expected_char =        
        consume_whitespace ();
        if !index < len && input_string.[!index] = expected_char then (
            incr index;
            true
        ) else 
            false
    in

    let parse_object () = 
        if not (consume_char '{') then false
        else if not (consume_char '}') then false
        else true
    in

    parse_object ()
