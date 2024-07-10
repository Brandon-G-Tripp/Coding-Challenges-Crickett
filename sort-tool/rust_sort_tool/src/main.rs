use std::env;
use std::io;
use std::io::Write;
use std::path::Path;

mod sort;
mod algorithms;

fn main() -> io::Result<()> {
    let mut args: Vec<String> = env::args().collect();
    let mut deduplicate = false;
    let mut algorithm = String::from("quick");

    let mut i = 1;
    while i < args.len() {
        match args[i].as_str() {
            "-u" => {
                deduplicate = true;
                args.remove(i);
            }
            "-a" => {
                if i + 1 < args.len() {
                    algorithm = args[i + 1].clone();
                    args.remove(i);
                    args.remove(i);
                } else {
                    eprintln!("Error: -a flag requires an algorithm name");
                    std::process::exit(1);
                }
            }
            _ => i += 1,
        }
    }

    if args.len() != 2 {
        eprintln!("Usage: {} [-u] [-a <algorithm>] <filename>", args[0]);
        std::process::exit(1);
    }

    let filename = &args[1];
    let path = Path::new(filename);
    let absolute_path = if path.is_relative() {
        env::current_dir()?.join(path)
    } else {
        path.to_path_buf()
    };

    let sorted_lines = sort::sort_file(absolute_path.to_str().unwrap(), deduplicate, &algorithm)?;

    let stdout = io::stdout();
    let mut handle = stdout.lock();
    for line in sorted_lines {
        handle.write_all(line.as_bytes())?;
        handle.write_all(b"\n")?;
    }

    Ok(())
}

#[cfg(test)]
mod tests {
    use std::process::Command;

    #[test]
    fn test_cli_with_words_file() -> Result<(), Box<dyn std::error::Error>> {
        let filename = "../words.txt";

        let output = Command::new("cargo")
            .args(&["run", "--", filename])
            .output()?;

        let cli_output = String::from_utf8(output.stdout)?;

        // Write the sorted lines to a temporary file
        let temp_file = tempfile::NamedTempFile::new()?;
        std::fs::write(temp_file.path(), cli_output)?;

        let piped_output = Command::new("uniq")
            .arg(temp_file.path())
            .output()?;

        let stdout = String::from_utf8(piped_output.stdout)?;
        let lines: Vec<&str> = stdout.lines().collect();

        // Check the first five lines of the sorted output
        let expected_lines = vec![
            "A",
            "ABACK",
            "ABANDON",
            "ABANDONED",
            "ABATED",
            "ABBREVIATED",
            "ABEYANCE",
            "ABIDE",
            "ABILITY",
            "ABLE",
        ];
        let actual_lines: Vec<_> = lines.iter().take(10).copied().collect();

        assert_eq!(actual_lines, expected_lines, "The first 5 lines do not match the expected output");

        Ok(())
    }

    #[test]
    fn test_cli_with_algorithm_flag() -> Result<(), Box<dyn std::error::Error>> {
        let filename = "test_words_algo_flag.txt";
        let content = "banana\napple\ncherry\ndate\n";
        std::fs::write(filename, content)?;

        for algorithm in &["merge", "quick", "heap", "radix"] {
            let output = Command::new("cargo")
                .args(&["run", "--", "-a", algorithm, filename])
                .output()?;

            let cli_output = String::from_utf8(output.stdout)?;
            let expected_output = "APPLE\nBANANA\nCHERRY\nDATE\n";

            assert_eq!(cli_output, expected_output, "Failed for algorithm: {}", algorithm);
        }

        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test] 
    fn test_cli_with_radix_sort_numbers() -> Result<(), Box<dyn std::error::Error>> {
        let filename = "test_numbers_radix.txt";
        let content = "10\n1\n100\n1000\n";
        std::fs::write(filename, content)?;

        let output = Command::new("cargo")
            .args(&["run", "--", "-a", "radix", filename])
            .output()?;

        let cli_output = String::from_utf8(output.stdout)?;
        let expected_output = "1\n10\n100\n1000\n";

        assert_eq!(cli_output, expected_output, "Failed for radix sort with numbers");

        std::fs::remove_file(filename)?;

        Ok(())
    }
}
