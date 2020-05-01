package intersection_of_two_arrays

// https://leetcode.com/problems/intersection-of-two-arrays

// time: O(n+m) 0ms 100%
// space: O(n+m) 3.2M

// leetcode submit region begin(Prohibit modification and deletion)

func intersection(nums1 []int, nums2 []int) (ret []int) {
	ret = []int{}
	m := map[int]int{}
	for _, v := range nums1 {
		m[v] = 1
	}
	for _, v := range nums2 {
		if m[v] == 1 {
			m[v]++
		}
	}
	for k, v := range m {
		if v == 2 {
			ret = append(ret, k)
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
