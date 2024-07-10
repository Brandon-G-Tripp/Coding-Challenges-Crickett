open OUnit2
open Ocaml_sort_tool.Sort

let test_sort_file _ = 
    let input = "banana\napple\ncherry\n" in
    let filename = "test_input.txt" in
    let oc = open_out filename in
    output_string oc input;
    close_out oc;
    let result = sort_file filename in
    let expected = ["apple"; "banana"; "cherry"] in
    assert_equal expected result ~printer:(String.concat ", ");
    Sys.remove filename

let test_sort_file_unique _ = 
    let input = "banana\napple\ncherry\napple\nbanana\n" in
    let filename = "test_input_unique.txt" in
    let oc = open_out filename in
    output_string oc input;
    close_out oc;
    let result = sort_file ~unique:true filename in
    let expected = ["apple"; "banana"; "cherry"] in
    assert_equal expected result ~printer:(String.concat ", ");
    Sys.remove filename

let suite = 
    "sort_tests" >::: [
        "test_sort_file" >:: test_sort_file;
        "test_sort_file_unique" >:: test_sort_file_unique;
    ]

let () = 
    run_test_tt_main suite
