open OUnit2
open Cut

let test_cut_second_field _ = 
    (* Prepare the input file *)
    let input_data = "a\tf1\t1\nb\t1\t2\nc\t6\t3\nd\t11\t4\ne\t16\t5\nf\t21\t6\n" in
    let oc = open_out "sample.tsv"in
    output_string oc input_data;
    close_out oc;

    (* Run the cut command and capture the output *) 
    let output = cut_second_field "sample.tsv" in

    let expected_output = "f1\n1\n6\n11\n16\n21" in
    assert_equal expected_output (String.concat "\n" output);

    Sys.remove "sample.tsv"

let suite = 
    "CutTest" >::: [
        "test_cut_second_field" >:: test_cut_second_field;
    ]

let _ = run_test_tt_main suite

