use std::fs::File;
use std::io::{BufRead, BufReader};
use crate::algorithms::{heap_sort, merge_sort, quick_sort, radix_sort};

pub fn sort_file(filename: &str, deduplicate: bool, algorithm: &str) -> std::io::Result<Vec<String>> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut lines: Vec<String> = reader
        .lines()
        .filter_map(|line| line.ok())
        .map(|line| line.trim().to_uppercase())
        .filter(|line| !line.is_empty())
        .collect();

    match algorithm {
        "merge" => merge_sort::sort(&mut lines), 
        "quick" => quick_sort::sort(&mut lines), 
        "heap" => heap_sort::sort(&mut lines), 
        "radix" => {
            // if all lines can be parsed as numbers, use radix sort
            if lines.iter().all(|s| s.parse::<u32>().is_ok()) {
                let mut numbers: Vec<u32> = lines
                    .iter()
                    .filter_map(|s| s.parse().ok())
                    .collect();
                radix_sort::sort(&mut numbers);
                lines = numbers.into_iter().map(|n| n.to_string()).collect();
            } else {
                // if not all lines are numbers, fall back to quick sort
                quick_sort::sort(&mut lines);
            }
        }
        _ => {
            eprintln!("Unknown sorting algorithm: {}. Using quick sort.", algorithm);
            quick_sort::sort(&mut lines);
        }
    }

    if deduplicate {
        lines.dedup();
    }

    lines.sort();

    Ok(lines)
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::env;

    // fn create_temp_file(name: &str, content: &str) -> Result<PathBuf, Box<dyn std::error::Error>> {
    //     let mut path = env::temp_dir();
    //     path.push(format!("{}_{}.txt", name, rand::random::<u32>()));
    //     let mut file = File::create(&path)?;
    //     file.write_all(content.as_bytes())?;
    //     Ok(path)
    // }

    #[test]
    fn test_sort_basic_functionality() -> Result<(), Box< dyn std::error::Error>> {
        // Create a temporary file with unsorted content
        let input = "banana\napple\ncherry\napple\nbanana\n";
        let filename = "test_input_basic.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename, false, "quick")?;
        let expected_lines = vec![
            "APPLE".to_string(),
            "APPLE".to_string(),
            "BANANA".to_string(),
            "BANANA".to_string(),
            "CHERRY".to_string()
        ];

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

        let sorted_lines = sort_file(filename, false, "quick")?;
        assert!(sorted_lines.is_empty());

        // Clean up 
        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_words_file() -> Result<(), Box<dyn std::error::Error>> {
        let filename = "words_test.txt";
        let sample_content =  "zebra\napple\nbanana\ncherry\ndate\n";
        std::fs::write(filename, sample_content)?;

        let sorted_lines = sort_file(filename, false, "quick")?;
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

    #[test] 
    fn test_sort_with_deduplication() -> Result<(), Box<dyn std::error::Error>> {
        let input = "apple\nbanana\ncherry\napple\nbanana\n";
        let filename = "test_input_with_dedup.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename, true, "quick")?;
        let expected_lines = vec!["APPLE".to_string(), "BANANA".to_string(), "CHERRY".to_string()];

        assert_eq!(sorted_lines, expected_lines);

        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_with_duplicates_and_deduplicate() -> Result<(), Box<dyn std::error::Error>> {
        let input = "banana\napple\ncherry\napple\nbanana\n";
        let filename = "test_input_without_dedup.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename, true, "quick")?;
        let expected_lines = vec!["APPLE".to_string(), "BANANA".to_string(), "CHERRY".to_string()];

        assert_eq!(sorted_lines, expected_lines);

        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_with_duplicates_and_no_deduplicate() -> Result<(), Box<dyn std::error::Error>> {
        let input = "banana\napple\ncherry\napple\nbanana\n";
        let filename = "test_input_with_dups.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename, false, "quick")?;
        let expected_lines = vec!["APPLE".to_string(), "APPLE".to_string(), "BANANA".to_string(), "BANANA".to_string(), "CHERRY".to_string()];

        assert_eq!(sorted_lines, expected_lines);

        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_file_uses_merge_sort() -> Result<(), Box<dyn std::error::Error>> {
        let input = "banana\napple\ncherry\ndate\nelderberry\n";
        let filename = "test_merge_sort_input.txt";
        std::fs::write(filename, input)?;

        let sorted_lines = sort_file(filename, false, "quick")?;
        let expected_lines = vec![
            "APPLE".to_string(),
            "BANANA".to_string(),
            "CHERRY".to_string(),
            "DATE".to_string(),
            "ELDERBERRY".to_string(),
        ];

        assert_eq!(sorted_lines, expected_lines);

        std::fs::remove_file(filename)?;

        Ok(())
    }

    #[test]
    fn test_sort_file_with_different_algorithms() -> std::io::Result<()> {
        let input = "banana\napple\ncherry\ndate\nelderberry\n";
        let filename = "test_algorithms.txt";
        std::fs::write(filename, input)?;

        let algorithms = vec!["merge", "quick", "heap"];
        let expected_lines = vec![
            "APPLE".to_string(),
            "BANANA".to_string(),
            "CHERRY".to_string(),
            "DATE".to_string(),
            "ELDERBERRY".to_string(),
        ];

        for algorithm in algorithms {
            let sorted_lines = sort_file(filename, false, algorithm)?;
            assert_eq!(sorted_lines, expected_lines, "Failed for algorithm: {}", algorithm);
        }

        std::fs::remove_file(filename)?;

        Ok(())
    }
}
