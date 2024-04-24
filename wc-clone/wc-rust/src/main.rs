use clap::Parser;
use std::fs::File;
use std::io::{BufRead, BufReader, Read};

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(short = 'c', long = "bytes")]
    count_bytes: bool, 

    #[arg(short = 'l', long = "lines")]
    count_lines: bool,

    #[arg(required = true)]
    file: String,
} 


fn main() {
    // Parse cli args
    let args = Args::parse();

    // Get the file path from args
    let file_path = &args.file;

    // count the number of bytes in the file
    if args.count_bytes {
        match count_bytes(file_path) {
            Ok(count) => println!("{} {}", count, file_path),
            Err(e) => {
                eprintln!("Error: {}", e);
                std::process::exit(1);
            } 
        } 
    } 

    // count the number of lines in the file
    if args.count_lines {
        match count_lines(file_path) {
            Ok(count) => println!("{} {}", count, file_path),
            Err(e) => {
                eprintln!("Error: {}", e);
                std::process::exit(1);
            } 
        } 
    } 

    if !args.count_bytes && !args.count_lines {
        eprintln!("Error: Missing -c or -l flag");
        std::process::exit(1);
    } 
}

fn count_lines(file_path: &str) -> Result<usize, std::io::Error> {
    let file = File::open(file_path)?;
    let reader = BufReader::new(file);
    let line_count = reader.lines().filter(|line| line.is_ok()).count();
    Ok(line_count)
}

fn count_bytes(file_path: &str) -> Result<usize, std::io::Error> {
    let mut file = File::open(file_path)?;
    let mut contents = Vec::new();
    file.read_to_end(&mut contents)?;
    Ok(contents.len())
} 

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs::File;
    use std::io::Write;

    #[test]
    fn test_count_bytes() {
        // Create a temporary file with sample content
        let mut file = File::create("test_count_bytes.txt").unwrap();
        file.write_all(b"Sample content").unwrap();

        // Call the count_bytes function
        let count = count_bytes("test_count_bytes.txt").unwrap();

        // Assert the expected byte count
        assert_eq!(count, 14);

        // clean up the temp file
        std::fs::remove_file("test_count_bytes.txt").unwrap();
    } 

    #[test]
    fn test_count_lines() {
        let mut file = File::create("test_count_lines.txt").unwrap();
        file.write_all(b"Line 1\nLine 2\nLine 3").unwrap();

        let count = count_lines("test_count_lines.txt").unwrap();

        assert_eq!(count, 3);

        std::fs::remove_file("test_count_lines.txt").unwrap();
    }

    #[test]
    fn test_file_not_found() {
        // Call the count_bytes function with a non-existent file
        let result = count_bytes("nonexistent.txt");

        // assert that the expected error is returned
        assert!(result.is_err());
    } 


    #[test]
    fn test_count_words() {
        let mut file = File::create("test_count_words.txt").unwrap();
        file.write_all(b"This is a sample file\n with multiple words\n on each line").unwrap();

        let count = count_words("test_count_words.txt").unwrap();

        assert_eq!(count, 12);

        std::fs::remove_file("test_count_words.txt").unwrap();
    } 
} 
