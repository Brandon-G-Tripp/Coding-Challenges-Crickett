let count_frequencies text = 
    let freq_map = Hashtbl.create 256 in
    String.iter (fun char -> 
        let count = Hashtbl.find_opt freq_map char |> Option.value ~default:0 in
        Hashtbl.replace freq_map char (count + 1)
    ) text;
    freq_map

let%test_unit "count_frequencies counts the frequencies of characters correctly" = 
    let text = "aabbccabc" in 
    let expected_frequencies = [('a', 3); ('b', 3); ('c', 3)] |> List.to_seq |> Hashtbl.of_seq in
    let actual_frequencies = count_frequencies text in
    Hashtbl.iter (fun char count -> 
        let expected_count = Hashtbl.find_opt expected_frequencies char |> Option.value ~default:0 in
        assert (count = expected_count)
    ) actual_frequencies

