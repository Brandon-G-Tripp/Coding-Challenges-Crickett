let count_bytes file_name = 
    let ic = open_in_bin file_name in
    let len = in_channel_length ic in
    close_in ic;
    len
