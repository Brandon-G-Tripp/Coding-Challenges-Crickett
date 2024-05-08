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

            if !is_valid_value(value) {
                return false;
            }
        }

        return true;
    } 

    false
}

fn is_valid_value(value: &str) -> bool {
    value.starts_with("\"") && value.ends_with("\"") && value.len() >= 2 ||
        value == "true" || value == "false" || value == "null" ||
        value.parse::<f64>().is_ok()
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

    #[test]
    fn test_valid_json_object_with_different_value_types() {
        let json = r#"{
            "key1": true,
            "key2": false,
            "key3": null,
            "key4": "value",
            "key5": 101
        }"#;

        assert_eq!(is_valid_json(json), true);
    }
} 
