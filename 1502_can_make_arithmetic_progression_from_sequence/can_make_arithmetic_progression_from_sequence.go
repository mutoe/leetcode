package can_make_arithmetic_progression_from_sequence

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence

// level: 1
// time: O(n*log(n)) 0ms 100%
// space: O(n) 2.6M 12.28%

// leetcode submit region begin(Prohibit modification and deletion)
func canMakeArithmeticProgression(arr []int) bool {
	tmp := make([]int, len(arr))
	mergeSort(arr, 0, len(arr)-1, tmp)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

func mergeSort(arr []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(arr, l, mid, tmp)
	mergeSort(arr, mid+1, r, tmp)

	i, j, k := l, mid+1, 0
	if arr[mid] <= arr[mid+1] {
		return
	}
	for ; i <= mid && j <= r; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; k, i = k+1, i+1 {
		tmp[k] = arr[i]
	}
	for ; j <= r; k, j = k+1, j+1 {
		tmp[k] = arr[j]
	}
	for k = 0; l <= r; l, k = l+1, k+1 {
		arr[l] = tmp[k]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
