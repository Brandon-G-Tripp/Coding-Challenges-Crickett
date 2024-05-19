mod huffman_tree;

use std::env;
use std::{collections::HashMap, fs::File};
use std::io::{Read};

use crate::huffman_tree::{Node, traverse_tree};

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Please provide a file path as a command line argument.");
        std::process::exit(1);
    }

    let file_path = &args[1];
    match read_file_contents(file_path) {
        Ok(contents) => {
            let frequencies = count_character_frequencies(&contents);
            println!("Character frequencies: {:?}", frequencies);

            let huffman_tree = Node::build_huffman_tree(&frequencies);
            println!("Huffman tree: {:?}", huffman_tree);
            traverse_tree(&huffman_tree, 0);
        }
        Err(e) => {
            eprintln!("Error reading file: {}", e);
        }
    }
}

fn read_file_contents(file_path: &str) -> Result<String, std::io::Error> {
    let mut file = File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}

pub fn count_character_frequencies(input: &str) -> HashMap<char, usize> {
    let mut frequencies = HashMap::new();

    for c in input.chars() {
        *frequencies.entry(c).or_insert(0) += 1;
    }
    frequencies
}



#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    use super::*;

    #[test]
    fn test_count_character_frequencies_from_file() {
        let file_path = "../LesMiserables.txt";
        match read_file_contents(file_path) {
            Ok(contents) => {
                let frequencies = count_character_frequencies(&contents);
                assert_eq!(333, *frequencies.get(&'X').unwrap_or(&0));
                assert_eq!(223000, *frequencies.get(&'t').unwrap_or(&0));
            }
            Err(e) => {
                eprintln!("Error reading file: {}", e);
                assert!(false, "File read error");
            }
        }
    }

    #[test]
    fn test_count_character_frequencies() {
        let input = "aabbcc";
        let expected_frequencies = HashMap::from([('a', 2), ('b', 2), ('c', 2)]);
        let actual_frequencies = count_character_frequencies(input);
        assert_eq!(expected_frequencies, actual_frequencies);
    }

    #[test]
    fn test_count_character_frequencies_with_special_chars() {
        let input = "Hello, World! \n\t";
        let expected_frequencies = HashMap::from([
            ('H', 1), ('e', 1), ('l', 3), ('o', 2), (',', 1), (' ', 2),
            ('W', 1), ('r', 1), ('d', 1), ('!', 1), ('\n', 1), ('\t', 1)
        ]);
        let actual_frequencies = count_character_frequencies(input);
        assert_eq!(expected_frequencies, actual_frequencies);
    }

    #[test]
    fn test_count_character_frequencies_unicode() {
        let input = "Hélló, Wørld!";
        let expected_frequencies = HashMap::from([
            ('H', 1), ('é', 1), ('l', 3), ('ó', 1), (',', 1), (' ', 1),
            ('W', 1), ('ø', 1), ('r', 1), ('d', 1), ('!', 1)
        ]);

        let actual_frequencies = count_character_frequencies(input);

        assert_eq!(expected_frequencies.len(), actual_frequencies.len());

        for (key, val) in expected_frequencies {
            assert_eq!(val, *actual_frequencies.get(&key).unwrap());
        }
    }
}
