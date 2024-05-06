pub fn is_valid_json(json: &str) -> bool {
    json.trim() == "{}"
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_valid_empty_object() {
        let json = "{}";
        assert_eq!(is_valid_json(json), true);
    }

    #[test]
    fn test_invalid_json() {
        let json = "{";
        assert_eq!(is_valid_json(json), false);
    }

    #[test]
    fn test_invalid_json_missing_closing_brace() {
        let json = "{";
        assert_eq!(is_valid_json(json), false);
    }

    #[test]
    fn test_invalid_json_missing_opening_brace() {
        let json = "}";
        assert_eq!(is_valid_json(json), false);
    } 
} 
