package insertionSort

func InsertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				continue
			}
		}
	}
	return arr
}
