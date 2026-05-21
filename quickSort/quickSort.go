package quickSort

func QuickSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, left int, right int) {
	if left < right {
		pivotIndex := partition(arr, left, right)
		sort(arr, left, pivotIndex-1)
		sort(arr, pivotIndex+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
