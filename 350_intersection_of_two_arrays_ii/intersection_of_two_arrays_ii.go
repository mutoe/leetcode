package intersection_of_two_arrays_ii

// https://leetcode.com/problems/intersection-of-two-arrays-ii

// level: 1
// time: O(n) 4ms 79.29%
// space: O(n) 3.1M 67.96%

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ret := make([]int, 0)
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num]--
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
