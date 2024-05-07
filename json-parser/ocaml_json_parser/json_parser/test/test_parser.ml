open OUnit2
open Json_parser.Parser

let test_parse_valid_json _ = 
    let valid_json = "{}" in
    assert_equal true (parse_json valid_json)

let test_parse_invalid_json _ = 
    let invalid_json = "invalid" in
    assert_equal false (parse_json invalid_json)

let test_parse_empty_object _ = 
    let json = "{}" in 
    assert_equal true (parse_json json)

let test_parse_single_pair_object _ = 
    let json = {|{"key": "value"}|} in
    assert_equal true (parse_json json)

let test_parse_multi_pair_object _ = 
    let json = {|{"key1": "value1", "key2": "value2"}|} in
    assert_equal true (parse_json json)

let test_parse_invalid_object _ = 
    let json = {|{"key": "value",}|} in
    assert_equal false (parse_json json)

let suite = 
    "Parser tests" >::: [
        "test_parse_valid_json" >:: test_parse_valid_json;
        "test_parse_invalid_json" >:: test_parse_invalid_json;
        "test_parse_empty_object" >:: test_parse_empty_object;
        "test_parse_single_pair_object" >:: test_parse_single_pair_object;
        "test_parse_multi_pair_object" >:: test_parse_multi_pair_object;
        "test_parse_invalid_object" >:: test_parse_invalid_object;
    ]

let _ = run_test_tt_main suite
