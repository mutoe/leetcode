package largest_perimeter_triangle

// https://leetcode.com/problems/largest-perimeter-triangle

// selection sort and early screening
// time: O(n*log(n)) 32ms 96.67%
// space: O(1) 6.4M

// leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(A []int) (ret int) {
	for i := 0; i < len(A); i++ {
		maxIndex := i
		for j := i + 1; j < len(A); j++ {
			if A[j] > A[maxIndex] {
				maxIndex = j
			}
		}
		A[i], A[maxIndex] = A[maxIndex], A[i]
		if i >= 2 && A[i]+A[i-1] > A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

// leetcode submit region end(Prohibit modification and deletion)
