package algorithms

func MergeSort(arr []string) {
    if len(arr) <= 1 {
        return 
    }

    mid := len(arr) / 2
    left := make([]string, mid)
    right := make([]string, len(arr)-mid)

    copy(left, arr[:mid])
    copy(right, arr[mid:])

    MergeSort(left)
    MergeSort(right)

    merge(arr, left, right)
}

func merge(arr, left, right []string) {
    i, j, k := 0, 0, 0

    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }

    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}
