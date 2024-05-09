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


    let parse_number () = 
        let start_index = !index in
        let is_valid = ref true in
        if consume_char '-' then ();
        if not (consume_char '0') then (
            if !index < len && input_string.[!index] >= '1' && input_string.[!index] <= '9' then
                while !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' do 
                    incr index
                done
            else
                is_valid := false
        );
        if !index < len && input_string.[!index] = '.' then (
            incr index;
            if !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' then
                while !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' do 
                    incr index
                done
            else
                is_valid := false
        );
        if !index < len && (input_string.[!index] = 'e' || input_string.[!index] = 'E') then (
            incr index;
            if !index < len && (input_string.[!index] = '+' || input_string.[!index] = '-') then
                incr index;
            if !index <len && input_string.[!index] >= '0' && input_string.[!index] <= '9' then
                while !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' do 
                    incr index
                done
            else
                is_valid := false
        );
        !is_valid && !index > start_index
    in

    let parse_value () = 
        consume_whitespace ();
        if !index < len then
            match input_string.[!index] with 
            | '"' -> parse_string ()
            | 't' -> consume_char 't' && consume_char 'r' && consume_char 'u' && consume_char 'e'
            | 'f' -> consume_char 'f' && consume_char 'a' && consume_char 'l' && consume_char 's' && consume_char 'e'
            | 'n' -> consume_char 'n' && consume_char 'u' && consume_char 'l' && consume_char 'l'
            | '-' | '0'..'9' -> parse_number ()
            | _ -> false
        else
            false
    in

    let parse_key_value_pair () = 
        if not (parse_string ()) then false
        else if not (consume_char ':') then false
        else parse_value ()
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



