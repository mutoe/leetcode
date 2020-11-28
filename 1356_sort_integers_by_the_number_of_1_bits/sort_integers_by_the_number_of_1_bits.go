package sort_integers_by_the_number_of_1_bits

import "sort"

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits

// level: 1
// time: O(n*log(n)) 12ms 13.46%
// space: O(n) 3.6M 100%

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

func sortByBits(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	sort.Slice(arr, func(i, j int) bool {
		a := countOf1Bit(arr[i])
		b := countOf1Bit(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})

	return arr
}

// leetcode submit region end(Prohibit modification and deletion)
