use clap::Parser;
use std::fs::File;
use std::io::{self, Write, BufRead, BufReader, Read};

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
struct Args {
    #[arg(short = 'c', long = "bytes")]
    count_bytes: bool, 

    #[arg(short = 'l', long = "lines")]
    count_lines: bool,

    #[arg(short = 'w', long = "words")]
    count_words: bool,

    #[arg(short = 'm', long = "chars")]
    count_chars: bool,

    // #[arg(required = true)]
    file: Option<String>,
} 


fn main() {
    // Parse cli args
    let args = Args::parse();
    let _ = main_with_args(args, &mut io::stdin(), &mut io::stdout(), &mut io::stderr());
}

fn main_with_args<R: Read, W: Write, E: Write>(args:Args, reader: &mut R, writer: &mut W, stderr: &mut E) -> Result<(), std::io::Error> {
    // Check if a filename is provided
    let file_path = match &args.file {
        Some(file) => file,
        None => "-",
    }; 

    let mut reader: Box<dyn Read> = if file_path == "-" {
        if atty::is(atty::Stream::Stdin) && !args.count_lines && !args.count_words && !args.count_bytes && !args.count_chars {
            print_usage(stderr)?;
            return Err(std::io::Error::new(std::io::ErrorKind::NotFound, "No input provided"));
        } 
        Box::new(reader)
    } else {
        match File::open(file_path) {
            Ok(file) => Box::new(file),
            Err(e) => {
                writeln!(stderr, "Error: {}: {}", file_path, e)?;
                print_usage(stderr)?;
                return Err(e);
            }
        }
    }; 


    if args.count_chars {
        match count_chars(&mut reader) {
            Ok(count) => writeln!(writer, "{} {}", count, file_path)?,
            Err(e) => {
                eprintln!("Error: {}", e);
                return Err(e);
            } 
        } 
    } 

    // count the number of bytes in the file
    if args.count_bytes {
        match count_bytes(&mut reader) {
            Ok(count) => writeln!(writer, "{} {}", count, file_path)?,
            Err(e) => {
                eprintln!("Error: {}", e);
                return Err(e);
            } 
        };
    }; 

    // count the number of lines in the file
    if args.count_lines {
        match count_lines(&mut reader) {
            Ok(count) => writeln!(writer, "{} {}", count, file_path)?,
            Err(e) => {
                eprintln!("Error: {}", e);
                return Err(e);
            } 
        }; 
    }; 

    if args.count_words {
        match count_words(&mut reader) {
            Ok(count) => writeln!(writer, "{} {}", count, file_path)?,
            Err(e) => {
                eprintln!("Error: {}", e);
                return Err(e);
            }
        };
    }; 

    if !args.count_bytes && !args.count_lines && !args.count_words && !args.count_chars {
        match default_count(&mut reader) {
            Ok((line_count, word_count, byte_count)) => {
                writeln!(writer, "{} {} {} {}", line_count, word_count, byte_count, file_path)?;
           }
            Err(e) => {
                eprintln!("Error: {}", e);
                return Err(e);
            }

        } 
    } 

    Ok(())
}

fn print_usage<W: Write>(writer: &mut W) -> std::io::Result<()> {
    writeln!(writer, "Usage: wc [OPTION]... [FILE]...")?;
    writeln!(writer, "Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified.")?;
    writeln!(writer, "With no FILE, or when FILE is -, read standard input.")?;
    writeln!(writer, "")?;
    writeln!(writer, "Options:")?;
    writeln!(writer, "  -c, --bytes     print the byte counts")?;
    writeln!(writer, "  -m, --chars     print the character counts")?;
    writeln!(writer, "  -l, --lines     print the newline counts")?;
    writeln!(writer, "  -w, --words     print the word counts")?;
    writeln!(writer, "")?;
    writeln!(writer, "If no option is specified, the default is to print the newline, word, and byte counts.")?;

    Ok(())
}

fn count_chars<R: Read>(mut reader: R) -> Result<usize, std::io::Error> {
    let mut contents = String::new();
    reader.read_to_string(&mut contents)?;
    Ok(contents.chars().count())
} 

fn count_words<R: Read>(reader: R) -> Result<usize, std::io::Error> {
    let reader = BufReader::new(reader);
    let word_count = reader
        .lines()
        .filter(|line| line.is_ok())
        .map(|line| line.unwrap().split_whitespace().count())
        .sum();
    Ok(word_count)
} 

fn count_lines<R: Read>(reader: R) -> Result<usize, std::io::Error> {
    let reader = BufReader::new(reader);
    let line_count = reader.lines().filter(|line| line.is_ok()).count();
    Ok(line_count)
}

fn count_bytes<R: Read>(mut reader: R) -> Result<usize, std::io::Error> {
    let mut contents = Vec::new();
    reader.read_to_end(&mut contents)?;
    Ok(contents.len())
} 

fn default_count<R: Read>(mut reader: R) -> Result<(usize, usize, usize), std::io::Error> {
    let mut buffer = Vec::new();
    reader.read_to_end(&mut buffer)?;

    let line_count = count_lines(buffer.as_slice())?;
    let word_count = count_words(buffer.as_slice())?;
    let byte_count = buffer.len();
    Ok((line_count, word_count, byte_count))
}

#[cfg(test)]
mod tests {

    use super::*;
    use std::io::Cursor;
    use std::fs::File;
    use std::io::Write;

    #[test]
    fn test_read_from_stdin() {
        let input = "Line 1\nLine 2\nLine 3\n";
        let mut stdin = Cursor::new(input.as_bytes());
        let mut stdout = Cursor::new(Vec::new());
        let mut stderr = Cursor::new(Vec::new());

        let args = Args {
            count_bytes: false, 
            count_lines: true, 
            count_words: false,
            count_chars: false, 
            file: None,
        };


        let _ = main_with_args(args, &mut stdin, &mut stdout, &mut stderr);

        let output_string = String::from_utf8(stdout.into_inner()).unwrap();
        assert_eq!(output_string, "3 -\n");
    }

    #[test]
    fn test_count_bytes() {
        // Create a temporary file with sample content
        let mut file = File::create("test_count_bytes.txt").unwrap();
        file.write_all(b"Sample content").unwrap();

        // Call the count_bytes function
        let count = count_bytes(&mut File::open("test_count_bytes.txt").unwrap()).unwrap();

        // Assert the expected byte count
        assert_eq!(count, 14);

        // clean up the temp file
        std::fs::remove_file("test_count_bytes.txt").unwrap();
    } 

    #[test]
    fn test_count_lines() {
        let mut file = File::create("test_count_lines.txt").unwrap();
        file.write_all(b"Line 1\nLine 2\nLine 3").unwrap();

        let count = count_lines(&mut File::open("test_count_lines.txt").unwrap()).unwrap();

        assert_eq!(count, 3);

        std::fs::remove_file("test_count_lines.txt").unwrap();
    }

    #[test]
    fn test_file_not_found() {
        let args = Args {
            count_bytes: true,
            count_lines: false,
            count_words: false,
            count_chars: false,
            file: Some("nonexistent.txt".to_string()),
        };

        let mut stdin = io::stdin();
        let mut stdout = Vec::new();
        let mut stderr = Vec::new();

        let result = main_with_args(args, &mut stdin, &mut stdout, &mut stderr);

        let output = String::from_utf8_lossy(&stderr);
        assert!(output.contains("Error: nonexistent.txt"));
        assert!(output.contains("No such file or directory"));
        assert!(result.is_err());
    } 


    #[test]
    fn test_count_words() {
        let mut file = File::create("test_count_words.txt").unwrap();
        file.write_all(b"This is a sample file\n with multiple words\n on each line").unwrap();

        let count = count_words(&mut File::open("test_count_words.txt").unwrap()).unwrap();

        assert_eq!(count, 11);

        std::fs::remove_file("test_count_words.txt").unwrap();
    } 

    #[test]
    fn test_count_chars() {
        let mut file = File::create("test_count_chars.txt").unwrap();
         file.write_all("Sample content with 🚀 emoji".as_bytes()).unwrap();

        let char_count = count_chars(&mut File::open("test_count_chars.txt").unwrap()).unwrap();
        let byte_count = count_bytes(&mut File::open("test_count_chars.txt").unwrap()).unwrap();

        assert_eq!(char_count, 27);
        assert_eq!(byte_count, 30);
        assert_ne!(char_count, byte_count);

        std::fs::remove_file("test_count_chars.txt").unwrap();
    }

    #[test]
    fn test_default_count() {
        let mut file = File::create("test_default_count.txt").unwrap();
        file.write_all(b"Line 1\nLine 2\nLine 3").unwrap();

        let (line_count, word_count, byte_count) = default_count(&mut File::open("test_default_count.txt").unwrap()).unwrap();

        assert_eq!(line_count, 3);
        assert_eq!(word_count, 6);
        assert_ne!(byte_count, 18);

        std::fs::remove_file("test_default_count.txt").unwrap();
    }
} 
