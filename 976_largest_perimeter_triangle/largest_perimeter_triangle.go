package largest_perimeter_triangle

import (
	"sort"
)

// https://leetcode.com/problems/largest-perimeter-triangle

// time: 40ms 90%
// space: 6.3M

// leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(A []int) (ret int) {
	aLen := len(A)
	sort.Ints(A)
	for i := aLen - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

// leetcode submit region end(Prohibit modification and deletion)
