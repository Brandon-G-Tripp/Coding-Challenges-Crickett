use std::env;
use std::fs;
use std::process;
use rust_json_parser::is_valid_json;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: {} <file_path>", args[0]);
        process::exit(1);
    } 

    let file_path = &args[1];
    let json = match fs::read_to_string(file_path) {
        Ok(content) => content, 
        Err(err) => {
            eprintln!("Error reading file: {}", err);
            process::exit(1);
        } 
    }; 

    if is_valid_json(&json) {
        println!("Valid JSON");
        process::exit(0);
    } else {
        println!("Invalid JSON");
        process::exit(1);
    }

}

#[cfg(test)]
mod test {
    use std::process::Command;

    #[test]
    fn test_invalid_json_file_step1() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg("tests/invalid.json")
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(1));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Invalid JSON");
    }

    #[test]
    fn test_valid_json_file_step1() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg("tests/valid.json")
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(0));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Valid JSON");
    }

    #[test]
    fn test_invalid_json_file_step2_case1() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg("tests/invalid.json")
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(1));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Invalid JSON");
    } 

    #[test]
    fn test_invalid_json_file_step2_case2() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg(concat!("tests/invalid2.json"))
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(1));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Invalid JSON");
    }

    #[test]
    fn test_valid_json_file_step2_case1() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg("tests/valid.json")
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(0));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Valid JSON");
    } 

    #[test]
    fn test_valid_json_file_step2_case2() {
        let output = Command::new("cargo")
            .arg("run")
            .arg("--")
            .arg("tests/valid2.json")
            .output()
            .expect("Failed to execute command");

        assert_eq!(output.status.code(), Some(0));
        assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "Valid JSON");
    }
}
