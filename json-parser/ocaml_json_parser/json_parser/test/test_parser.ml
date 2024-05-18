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

let test_parse_string_value _ = 
    let json = {|{"key": "value"}|} in
    assert_equal true (parse_json json)
    
let test_parse_numeric_value _ = 
    let json = {|{"key": 42}|} in
    assert_equal true (parse_json json)

let test_parse_boolean_value  _ = 
    let json = {|{"key": true}|} in
    assert_equal true (parse_json json)

let test_parse_null_value _ = 
    let json = {|{"key": null}|} in
    assert_equal true (parse_json json)

let test_parse_all_value_types _ =
    let json = {|{"key1": "string", "key2": 42, "key3": true, "key4": false, "key5": null}|} in
    assert_equal true (parse_json json)

let test_parse_nested_object _ =
    let json = {|{"key": {"nestedKey": "nestValue"}}|} in
    assert_equal true (parse_json json)

let test_parse_nested_array _ = 
    let json = {|{"key": ["value1", "value2"]}|} in
    assert_equal true (parse_json json)

let test_parse_array_of_objects _ = 
    let json = {|{"key": [{"nestedKey1": "nestedValue1"}, {"nestedKey2": "nestedValue2"}]}|} in
    assert_equal true (parse_json json)

let test_parse_complex_object _ = 
    let json = {|{
        "key1": "value1",
        "key2": 42,
        "key3": true,
        "key4": false,
        "key5": null,
        "key6": {"nestedKey1": "nestedValue1", "nestedKey2": [1, 2, 3]},
        "key7": [{"nestedKey3": "nestedValue3"}, {"nestedKey4": "nestedValue4"}]
    }|} in
    assert_equal true (parse_json json)

let suite = 
    "Parser tests" >::: [
        "test_parse_valid_json" >:: test_parse_valid_json;
        "test_parse_invalid_json" >:: test_parse_invalid_json;
        "test_parse_empty_object" >:: test_parse_empty_object;
        "test_parse_single_pair_object" >:: test_parse_single_pair_object;
        "test_parse_multi_pair_object" >:: test_parse_multi_pair_object;
        "test_parse_invalid_object" >:: test_parse_invalid_object;
        "test_parse_string_value" >:: test_parse_string_value;
        "test_parse_numeric_value" >:: test_parse_numeric_value;
        "test_parse_boolean_value" >:: test_parse_boolean_value;
        "test_parse_null_value" >:: test_parse_null_value;
        "test_parse_nested_object" >:: test_parse_nested_object;
        "test_parse_nested_array" >:: test_parse_nested_array;
        "test_parse_array_of_objects" >:: test_parse_array_of_objects;
        "test_parse_complex_object" >:: test_parse_complex_object;
        "test_parse_all_value_types" >:: test_parse_all_value_types;
    ]

let _ = run_test_tt_main suite
