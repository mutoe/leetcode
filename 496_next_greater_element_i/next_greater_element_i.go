package next_greater_element_i

// https://leetcode.com/problems/next-greater-element-i

// level: 1
// time: O(n^2) 0ms 100%
// space: O(n) 3.4M 60.92%

// TODO: Optimize to O(n) (using Stack)

// leetcode submit region begin(Prohibit modification and deletion)
func findNext(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return i
		}
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1))
	greaterMap := map[int]int{}

	for i := 0; i < len(nums2); i++ {
		greaterMap[nums2[i]] = i + 1
	}

	for i := 0; i < len(nums1); i++ {
		index := greaterMap[nums1[i]]
		if index > 0 {
			el := findNext(nums2[index:], nums1[i])
			if el >= 0 {
				ret[i] = nums2[el+index]
				continue
			}
		}
		ret[i] = -1
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
