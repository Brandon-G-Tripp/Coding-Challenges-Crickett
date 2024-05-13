pub fn is_valid_json(json: &str) -> bool {
    let json = json.trim();

    if json.starts_with("{") && json.ends_with("}") {
        let content = &json[1..json.len() - 1].trim();
        if content.is_empty() {
            return true;
        }

        let mut pairs = Vec::new();
        let mut start = 0;
        let mut depth = 0;

        for (i, c) in content.char_indices() {
            match c {
                '{' | '[' => depth += 1,
                '}' | ']' => depth -= 1,
                ',' if depth == 0 => {
                    let pair = content[start..i].trim().to_string();
                    if !is_valid_key_value_pair(&pair) {
                        return false;
                    } 
                    pairs.push(pair);
                    start = i + 1;
                }
                _ => (),
            }
        }

        if depth != 0 {
            return false;
        }

        if start < content.len() {
            let remaining_pair = content[start..].trim().to_string();
            if !is_valid_key_value_pair(&remaining_pair) {
                return false;
            }
            pairs.push(remaining_pair);
        }

        if content.ends_with(",") {
            return false;
        }

        return true;
    } else if json.starts_with("[") && json.ends_with("]") {
        return is_valid_array(json);
    } else if is_valid_string(json) || is_valid_number(json) || is_valid_boolean(json) || json == "null" {
        return true;
    }

    false
}

fn is_valid_number(num: &str) -> bool {
    num.parse::<f64>().is_ok()
} 

fn is_valid_boolean(value: &str) -> bool {
    value == "true" || value == "false"
}

fn is_valid_key_value_pair(pair: &str) -> bool {
    let parts: Vec<&str> = pair.splitn(2, ":").map(|s| s.trim()).collect();
    if parts.len() != 2 {
        return false;
    }

    let key = parts[0].trim();
    let value = parts[1].trim();

    if !is_valid_string(key) {
        return false;
    }

    if !is_valid_value(value) {
        return false;
    }

    true
}

fn is_valid_value(value: &str) -> bool {
    is_valid_string(value)
        || value == "true"
        || value == "false"
        || value == "null"
        || value.parse::<f64>().is_ok()
        || is_valid_object(value)
        || is_valid_array(value)
}

fn is_valid_string(s: &str) -> bool {
    s.starts_with("\"") && s.ends_with("\"") && s.len() >= 2
}

fn is_valid_object(obj: &str) -> bool {
    obj.starts_with("{") && obj.ends_with("}") && {
        let content = &obj[1..obj.len() - 1].trim();
        if content.is_empty() {
            return true
        }
        is_valid_json(content)
    }
}

fn is_valid_array(arr: &str) -> bool {
    arr.starts_with("[") && arr.ends_with("]") && {
        let content = &arr[1..arr.len() - 1].trim();
        if content.is_empty() {
            return true;
        } 
        let values: Vec<&str> = content.split(",").map(|s| s.trim()).collect();
        values.iter().all(|&value| is_valid_json(value))
    }
}

fn is_valid_array_content(arr: &str) -> bool {
    let content = &arr[1..arr.len() - 1].trim();
    if content.is_empty() {
        return true;
    }

    let values: Vec<&str> = content.split(",").map(|s| s.trim()).collect();
    for value in values {
        if !is_valid_value(value) {
            return false;
        }
    }

    true
}
