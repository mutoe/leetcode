package relative_sort_array

// https://leetcode.com/problems/relative-sort-array

// level: 1
// time: O(n^2) 0ms 100%
// space: O(n) 2.3M 95.24%

// leetcode submit region begin(Prohibit modification and deletion)
func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, q[l]
	for i < j {
		i++
		j--
		for q[i] < x {
			i++
		}
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	i2, i1 := 0, 0
	for ; i2 < len(arr2); i2++ {
		for i1 < len(arr1)-1 && arr1[i1] == arr2[i2] {
			i1++
		}
		for j := i1 + 1; j < len(arr1); j++ {
			if arr1[j] == arr2[i2] {
				arr1[j], arr1[i1] = arr1[i1], arr1[j]
				i1++
			}
		}
	}
	if i1 < len(arr1) {
		quickSort(arr1, i1, len(arr1)-1)
	}
	return arr1
}

// leetcode submit region end(Prohibit modification and deletion)
