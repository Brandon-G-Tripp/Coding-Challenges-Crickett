package algorithms

func RadixSort(arr []string) {
    if len(arr) <= 1 {
        return
    }

    // Find the maximum length string
    maxLen := len(arr[0])
    for _, s := range arr {
        if len(s) > maxLen {
            maxLen = len(s)
        }
    }

    // Perform counting sort for every character from right to left
    for i := maxLen - 1; i >= 0; i-- {
        countingSort(arr,i)
    }
}

func countingSort(arr []string, position int) {
    n := len(arr)
    output := make([]string, n)
    count := make([]int, 256) // assuming ascii chars

    // Store count of occurrences
    for _, s := range arr {
        index := 0
        if position < len(s) {
            index = int(s[position])
        }
        count[index]++
    }

    // Change count[i] so that count[i] now contains actual
    // position of this character in output array
    for i := 1; i < 256; i++ {
        count[i] += count[i-1]
    }

    // Build the output array
    for i := n - 1; i >= 0; i-- {
        index := 0
        if position < len(arr[i]) {
            index = int(arr[i][position])
        }
        output[count[index]-1] = arr[i]
        count[index]--
    }

    // Copy the output array to arr[], so that arr[] now
    // contains sorted characters
    for i := 0; i < n; i++ {
        arr[i] = output[i]
    }
}
