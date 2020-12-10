package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j, x := l-1, r+1, arr[l]
	for i < j {
		for ok := true; ok; ok = arr[i] < x {
			i++
		}
		for ok := true; ok; ok = arr[j] > x {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}

func quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	arr[l], arr[mid] = arr[mid], arr[l]

	// arr[l+1, lt] < v
	// arr[lt+1, i-1] == v
	// arr[gt, r] > v
	i, lt, gt := l+1, l, r+1
	for i < gt {
		if arr[i] == arr[l] {
			i++
		} else if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else {
			gt--
			arr[gt], arr[i] = arr[i], arr[gt]
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]

	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}
