package sort_array_by_parity_ii

// https://leetcode.com/problems/sort-array-by-parity-ii

// level: 1
// time: O(n) 16ms 98.46%
// space: O(1) 6.1M 70.77%

// leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParityII(A []int) []int {
	odd, even := 0, 1
	l := len(A)
	for odd < l && even < l {
		if odd < l && A[odd]%2 == 0 {
			odd += 2
			continue
		}
		if even < l && A[even]%2 == 1 {
			even += 2
			continue
		}
		A[odd], A[even] = A[even], A[odd]
		odd += 2
		even += 2
	}

	return A
}

// leetcode submit region end(Prohibit modification and deletion)
