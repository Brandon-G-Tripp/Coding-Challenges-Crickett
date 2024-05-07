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

    let parse_string () =
        if not (consume_char '"') then false
        else (
            while !index < len && input_string.[!index] <> '"' do 
                if input_string.[!index] = '\\' then incr index;
                incr index
            done;
            consume_char '"'
        )
    in

    let parse_key_value_pair () = 
        if not (parse_string ()) then false
        else if not (consume_char ':') then false
        else parse_string ()
    in

    let rec parse_pairs () = 
        consume_whitespace ();
        if consume_char '}' then true
        else if not (parse_key_value_pair ()) then false
        else (
            consume_whitespace ();
            if consume_char ',' then (
                consume_whitespace ();
                if consume_char '}' then false
                else parse_pairs ()
            ) else
                consume_char '}'
        )
    in

    if not (consume_char '{') then false 
    else (
        if parse_pairs () then (
            consume_whitespace ();
            !index = len
        ) else 
            false
    )

