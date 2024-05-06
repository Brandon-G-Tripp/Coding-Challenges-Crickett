open OUnit2
open Json_parser.Parser

let test_parse_valid_json _ = 
    let valid_json = "{}" in
    assert_equal true (parse_json valid_json)

let test_parse_invalid_json _ = 
    let invalid_json = "invalid" in
    assert_equal false (parse_json invalid_json)


let suite = 
    "Parser tests" >::: [
        "test_parse_valid_json" >:: test_parse_valid_json;
        "test_parse_invalid_json" >:: test_parse_invalid_json;
    ]

let _ = run_test_tt_main suite
