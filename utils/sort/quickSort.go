package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l-1, r+1
	for i < j {
		for ok := true; ok; ok = arr[i] < arr[l] {
			i++
		}
		for ok := true; ok; ok = arr[j] > arr[l] {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i, j = i+1, j-1
		}
	}

	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}
