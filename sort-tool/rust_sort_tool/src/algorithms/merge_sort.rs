pub fn sort(arr: &mut [String]) {
    if arr.len() > 1 {
        let mid = arr.len() / 2;
        sort(&mut arr[..mid]);
        sort(&mut arr[mid..]);
        merge(arr, mid);
    }
}

fn merge(arr: &mut [String], mid: usize) {
    let left = arr[..mid].to_vec();
    let right = arr[mid..].to_vec();

    let mut i = 0;
    let mut j = 0;
    let mut k = 0;

    while i < left.len() && j < right.len() {
        if left[i] <= right[j] {
            arr[k] = left[i].clone();
            i += 1;
        } else {
            arr[k] = right[j].clone();
            j += 1;
        }
        k += 1;
    }

    while i < left.len() {
        arr[k] = left[i].clone();
        i += 1;
        k += 1;
    }

    while j < right.len() {
        arr[k] = right[j].clone();
        j += 1;
        k += 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test] 
    fn test_merge_sort() {
        let mut input = vec![
            "banana".to_string(),
            "apple".to_string(),
            "cherry".to_string(),
            "date".to_string(),
            "elderberry".to_string(),
        ];

        let expected = vec![
            "apple".to_string(),
            "banana".to_string(),
            "cherry".to_string(),
            "date".to_string(),
            "elderberry".to_string(),
        ];

        sort(&mut input);

        assert_eq!(input, expected);
    }
}
