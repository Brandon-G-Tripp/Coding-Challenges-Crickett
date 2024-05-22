use std::{env, process};

use cut::cut_second_field;

mod cut;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 || args[1] != "-f2" {
        eprintln!("Usage: cut -f2 <file>");
        process::exit(1);
    }

    let file_path = &args[2];
    match cut_second_field(file_path) {
        Ok(output) => {
            for line in output {
                println!("{}", line);
            }
        }
        Err(err) => {
            eprintln!("Error: {}", err);
            process::exit(1);
        }
    }
}

#[cfg(test)]
mod tests {
    use std::fs::File;
    use std::io::{BufRead, BufReader};

    #[test]
    fn test_cut_second_field() {
        let file_path = "../sample.tsv";
        let expected_output = vec![
            "f1",
            "1",
            "6",
            "11",
            "16",
            "21",
        ];

        let file = File::open(file_path).expect("Failed to open file");
        let reader = BufReader::new(file);

        let mut actual_output = Vec::new();
        for line in reader.lines() {
            let line = line.expect("Failed to read line");
            let fields: Vec<&str> = line.split('\t').collect();
            if fields.len() >= 2 {
                actual_output.push(fields[1].to_string());
            }
        }

        assert_eq!(actual_output, expected_output);
    }
}
