package mergeSort

func MergeSort(arr []int) []int {
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, left int, right int) {
	if left < right {
		mid := left + (right-left)/2
		sort(arr, left, mid)
		sort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	l := make([]int, n1)
	r := make([]int, n2)

	for i := 0; i < n1; i++ {
		l[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = arr[mid+1+j]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if l[i] <= r[j] {
			arr[k] = l[i]
			i++
		} else {
			arr[k] = r[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = l[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = r[j]
		j++
		k++
	}
}
