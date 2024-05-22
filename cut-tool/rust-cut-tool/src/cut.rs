use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn cut_second_field(file_path: &str) -> Result<Vec<String>, String> {
    let file = File::open(file_path).expect("Failed to open file");
    let reader = BufReader::new(file);

    let mut output = Vec::new();
    for line in reader.lines() {
        let line = line.expect("Failed to read line");
        let fields: Vec<&str> = line.split('\t').collect();
        if fields.len() >= 2 {
            output.push(fields[1].to_string());
        }
    }

    Ok(output)
}
