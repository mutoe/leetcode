package sort_integers_by_the_number_of_1_bits

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits

// level: 1
// time: O(n*log(n)) 60ms 13.46%
// space: O(n) 9.8M 11.54%

// leetcode submit region begin(Prohibit modification and deletion)

func countOf1Bit(x int) (count int) {
	if x == 0 {
		return
	}
	for ; x > 0; x >>= 1 {
		if x&1 == 1 {
			count++
		}
	}
	return
}

func mergeSort(q, cq []int, l, r int) {
	if l >= r || len(q) != len(cq) {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, cq, l, mid)
	mergeSort(q, cq, mid+1, r)

	tmp := make([]int, len(q))
	ctmp := make([]int, len(q))
	i, j, k := l, mid+1, 0
	for ; i <= mid && j <= r; k++ {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			ctmp[k] = cq[i]
			i++
		} else {
			tmp[k] = q[j]
			ctmp[k] = cq[j]
			j++
		}
	}
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = q[i]
		ctmp[k] = cq[i]
	}
	for ; j <= r; j, k = j+1, k+1 {
		tmp[k] = q[j]
		ctmp[k] = cq[j]
	}

	i, k = l, 0
	for ; i <= r; i, k = i+1, k+1 {
		q[i] = tmp[k]
		cq[i] = ctmp[k]
	}
}

func mapToBit(arr []int) []int {
	ret := make([]int, len(arr))
	for i, num := range arr {
		ret[i] = countOf1Bit(num)
	}
	return ret
}

func sortByBits(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	bitArr := mapToBit(arr)
	mergeSort(arr, bitArr, 0, len(arr)-1)
	mergeSort(bitArr, arr, 0, len(arr)-1)

	return arr
}

// leetcode submit region end(Prohibit modification and deletion)
