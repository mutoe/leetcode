package next_greater_element_i

// https://leetcode.com/problems/next-greater-element-i

// level: 1
// time: O(n^2) 4ms 86.21%
// space: O(n) 2.8M 91.95%

// leetcode submit region begin(Prohibit modification and deletion)
func find(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1))
Outer:
	for i := 0; i < len(nums1); i++ {
		targetIndex := find(nums2, nums1[i])
		if targetIndex < 0 {
			ret[i] = -1
			continue
		}
		rest := nums2[targetIndex+1:]
		for j := 0; j < len(rest); j++ {
			if nums2[j+targetIndex+1] > nums2[targetIndex] {
				ret[i] = nums2[j+targetIndex+1]
				continue Outer
			}
		}
		ret[i] = -1
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
