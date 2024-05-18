let parse_json input_string =
    Printf.printf "Parsing JSON: %s\n" input_string;
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
            if !index < len then (
                incr index;
                true
            ) else
                false
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
            if !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' then
                while !index < len && input_string.[!index] >= '0' && input_string.[!index] <= '9' do
                    incr index
                done
            else
                is_valid := false
        );
        !is_valid && !index > start_index
    in

    let rec parse_array () =
        Printf.printf "parsing value at index %d\n" !index;
        consume_whitespace ();
        if consume_char '[' then (
            consume_whitespace ();
            if consume_char ']' then true
            else (
                let rec parse_array_elements () =
                    let success = parse_value () in
                    consume_whitespace ();
                    if success then (
                        if consume_char ',' then (
                            consume_whitespace ();
                            parse_array_elements ()
                        ) else (
                            consume_char ']'
                        )
                    ) else
                        false
                in
                parse_array_elements ()
            )
        ) else
            false

    and parse_value () =
        Printf.printf "parsing value at index %d\n" !index;
        consume_whitespace ();
        if !index < len then
            match input_string.[!index] with
            | '"' -> parse_string ()
            | 't' -> consume_char 't' && consume_char 'r' && consume_char 'u' && consume_char 'e'
            | 'f' -> consume_char 'f' && consume_char 'a' && consume_char 'l' && consume_char 's' && consume_char 'e'
            | 'n' -> consume_char 'n' && consume_char 'u' && consume_char 'l' && consume_char 'l'
            | '-' | '0'..'9' -> parse_number ()
            | '[' -> parse_array ()
            | '{' -> parse_object ()
            | _ -> false
        else
            false

    and parse_key_value_pair () =
        if not (parse_string ()) then false
        else (
            consume_whitespace ();
            if not (consume_char ':') then false
            else (
                consume_whitespace ();
                parse_value ()
            )
        )

    and parse_object () =
        Printf.printf "parsing value at index %d\n" !index;
        consume_whitespace ();
        if consume_char '{' then (
            consume_whitespace ();
            if consume_char '}' then true
            else (
                let rec parse_object_elements () =
                    if not (parse_key_value_pair ()) then false
                    else (
                        consume_whitespace ();
                        if consume_char ',' then (
                            consume_whitespace ();
                            parse_object_elements ()
                        ) else (
                            consume_char '}'
                        )
                    )
                in
                parse_object_elements ()
            )
        ) else
            false
    in

    let result = parse_value () in
    consume_whitespace ();
    if result && !index = len then true
    else (
        consume_whitespace ();
        !index = len
    )
