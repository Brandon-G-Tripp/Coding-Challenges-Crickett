use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn sort_file(filename: &str) -> std::io::Result<Vec<String>> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut lines: Vec<String> = reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| line.trim().to_uppercase())
        .filter(|line| !line.is_empty())
        .collect();

    lines.sort();

    Ok(lines)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sort_basic_functionality() -> Result<(), Box< dyn std::error::Error>> {
        // Create a temporary file with unsorted content
        let input = "banana\napple\ncherry\n";
        let filename = "test_input.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename)?;
        let expected_lines = vec!["APPLE".to_string(), "BANANA".to_string(), "CHERRY".to_string()];

        assert_eq!(sorted_lines, expected_lines);

        // Clean up 
        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_empty_file() -> Result<(), Box<dyn std::error::Error>> {
        // Create an empty file 
        let filename = "empty.txt";
        std::fs::write(filename, "")?;

        let sorted_lines = sort_file(filename)?;
        assert!(sorted_lines.is_empty());

        // Clean up 
        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_words_file() -> Result<(), Box<dyn std::error::Error>> {
        let filename = "words.txt";
        let sample_content =  "zebra\napple\nbanana\ncherry\ndate\n";
        std::fs::write(filename, sample_content)?;

        let sorted_lines = sort_file(filename)?;
        assert_eq!(
            sorted_lines,
            vec![
            "APPLE".to_string(),
            "BANANA".to_string(),
            "CHERRY".to_string(),
            "DATE".to_string(),
            "ZEBRA".to_string()
            ]
        );

        std::fs::remove_file(filename)?;

        Ok(())
    }
}
