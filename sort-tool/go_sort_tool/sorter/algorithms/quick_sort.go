package algorithms

func QuickSort(arr []string) {
    if len(arr) <= 1 {
        return 
    }
    quickSortHelper(arr, 0, len(arr) - 1)
}

func quickSortHelper(arr []string, low, high int) {
    if low < high {
        pivotIndex := partition(arr, low, high)
        quickSortHelper(arr, low, pivotIndex-1)
        quickSortHelper(arr, pivotIndex + 1, high)
    }
}

func partition(arr []string, low, high int) int {
    pivot := arr[high]
    i := low - 1

    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
