package sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)

	i, j, k := l, mid+1, 0
	tmp := make([]int, r-l+1)
	for ; i <= mid && j <= r; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = arr[i]
	}
	for ; j <= r; j, k = j+1, k+1 {
		tmp[k] = arr[j]
	}
	for i, k = l, 0; i <= r; i, k = i+1, k+1 {
		arr[i] = tmp[k]
	}
}
