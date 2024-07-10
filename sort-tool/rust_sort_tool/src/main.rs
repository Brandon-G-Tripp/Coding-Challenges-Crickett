use std::env;
use std::io;
use std::io::Write;
use std::path::Path;

mod sort;
mod algorithms;

fn main() -> io::Result<()> {
    let mut args: Vec<String> = env::args().collect();
    let mut deduplicate = false;

    if args.len() > 1 && args[1] == "-u" {
        deduplicate = true;
        args.remove(1); // Remove the "-u" option from args
    }

    if args.len() != 2 {
        eprintln!("Usage: {} [-u] <filename>", args[0]);
        std::process::exit(1);
    }

    let filename = &args[1];
    let path = Path::new(filename);
    let absolute_path = if path.is_relative() {
        env::current_dir()?.join(path)
    } else {
        path.to_path_buf()
    };

    let sorted_lines = sort::sort_file(absolute_path.to_str().unwrap(), deduplicate)?;

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
}
