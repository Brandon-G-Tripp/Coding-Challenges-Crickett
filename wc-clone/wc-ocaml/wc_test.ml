open OUnit2
open Wc

let test_count_bytes _ =
    let content = "Sample content" in 
    let file_name = "test.txt" in 
    let oc = open_out file_name in
    output_string oc content;
    close_out oc;
    let count = count_bytes file_name in 
    assert_equal (String.length content) count;
    Sys.remove file_name

let test_count_bytes_file_not_found _ = 
    let file_name = "nonexistent.txt" in 
    try 
        let _ = count_bytes file_name in
        assert_failure "Expected an exception, but none was raised"
    with Sys_error _ -> ()

let test_count_lines _ =
    let content = "Line 1\nLine 2\nLine 3\n" in 
    let file_name = "test.txt" in
    let oc = open_out file_name in
    output_string oc content;
    close_out oc;
    let count = count_lines file_name in 
    assert_equal 3 count;
    Sys.remove file_name

let test_count_lines_file_not_found _ = 
    let file_name = "nonexistent.txt" in 
    try 
        let _ = count_lines file_name in
        assert_failure "Expected an exception, but none was raised"
    with Sys_error _ -> ()

let suite = 
    "CountBytes tests" >::: [
        "test_count_bytes" >:: test_count_bytes;
        "test_count_bytes_file_not_found" >:: test_count_bytes_file_not_found;
        "test_count_lines" >:: test_count_lines;
        "test_count_lines_file_not_found" >:: test_count_lines_file_not_found;
    ]

let () = 
    run_test_tt_main suite

