pub fn is_valid_json(json: &str) -> bool {
    let json = json.trim();

    if json.starts_with("{") && json.ends_with("}") {
        let content = &json[1..json.len() - 1].trim();
        if content.is_empty() {
            return true;
        }

        let pairs: Vec<&str> = content.split(",").map(|s| s.trim()).collect();
        for pair in pairs {
            let parts: Vec<&str> = pair.splitn(2, ":").map(|s| s.trim()).collect();
            if parts.len() != 2{
                return false;
            }

            let key = parts[0].trim();
            let value = parts[1].trim();

            if !key.starts_with("\"") || !key.ends_with("\"") || key.len() < 2 {
                return false;
            }

            if !value.starts_with("\"") || !value.ends_with("\"") || value.len() < 2 {
                return false;
            }
        }

        return true;
    } 

    false
}

#[cfg(test)]
mod test {
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
