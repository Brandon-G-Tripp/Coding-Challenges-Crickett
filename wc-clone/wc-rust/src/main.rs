use clap::Parser;
use std::fs::File;
use std::io::Read;

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(short = 'c', long = "bytes")]
    count_bytes: bool, 

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
    } else {
        eprintln!("Error: Missing -c flag");
        std::process::exit(1);
    } 
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
    use std::path::PathBuf;

    #[test]
    fn test_count_bytes() {
        // Create a temporary file with sample content
        let mut file = File::create("test.txt").unwrap();
        file.write_all(b"Sample content").unwrap();

        // Call the count_bytes function
        let count = count_bytes("test.txt").unwrap();

        // Assert the expected byte count
        assert_eq!(count, 14);

        // clean up the temp file
        std::fs::remove_file("test.txt").unwrap();
    } 

    #[test]
    fn test_file_not_found() {
        // Call the count_bytes function with a non-existent file
        let result = count_bytes("nonexistent.txt");

        // assert that the expected error is returned
        assert!(result.is_err());
    } 
} 
