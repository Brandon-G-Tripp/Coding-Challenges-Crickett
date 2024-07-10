pub fn sort(arr: &mut [u32]) {
    if arr.is_empty() {
        return;
    }
    
    let max = *arr.iter().max().unwrap();
    let mut exp = 1;

    while max / exp > 0 {
        counting_sort(arr, exp);
        exp *= 10;
    }
}

fn counting_sort(arr: &mut [u32], exp: u32) {
    let mut output = vec![0; arr.len()];
    let mut count = vec![0; 10];

    // Store count of occurrences in count[]
    for &num in arr.iter() {
        count[((num / exp) % 10) as usize] += 1;
    }

    // Change count[i] so that count[i] now contains
    // actual position of this digit in output[]
    for i in 1..10 {
        count[i] += count[i - 1];
    }

    // build the output array
    for &num in arr.iter().rev() {
        let index = ((num / exp) % 10) as usize;
        output[count[index] - 1] = num;
        count[index] -= 1;
    }

    // copy the output array to arr[], so that arr[] now
    // contains sorted numbers according to current digit
    arr.copy_from_slice(&output);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_radix_sort_empty() {
        let mut arr: Vec<u32> = vec![];
        sort(&mut arr);
        assert_eq!(arr, vec![]);
    }

    #[test]
    fn test_radix_sort_single_element() {
        let mut arr = vec![1];
        sort(&mut arr);
        assert_eq!(arr, vec![1]);
    }

    #[test]
    fn test_radix_sort_sorted_array() {
        let mut arr = vec![1, 2, 3, 4, 5];
        sort(&mut arr);
        assert_eq!(arr, vec![1, 2, 3, 4, 5]);
    }

    #[test]
    fn test_radix_sort_reverse_sorted_array() {
        let mut arr = vec![5, 4, 3, 2, 1];
        sort(&mut arr);
        assert_eq!(arr, vec![1, 2, 3, 4, 5]);
    }

    #[test] 
    fn test_radix_sort_unsorted_array() {
        let mut arr = vec![170, 45, 75, 90, 802, 24, 2, 66];
        sort(&mut arr);
        assert_eq!(arr, vec![2, 24, 45, 66, 75, 90, 170, 802]);
    }

    #[test]
    fn test_radix_sort_with_zeros() {
        let mut arr = vec![0, 10, 1, 100, 1000, 10000];
        sort(&mut arr);
        assert_eq!(arr, vec![0, 1, 10, 100, 1000, 10000]);
    }

    #[test]
    fn test_radix_sort_large_numbers() {
        let mut arr = vec![1_000_000, 10_000, 100_000, 1_000, 10_000_000];
        sort(&mut arr);
        assert_eq!(arr, vec![1_000, 10_000, 100_000, 1_000_000, 10_000_000]);
    }
}
