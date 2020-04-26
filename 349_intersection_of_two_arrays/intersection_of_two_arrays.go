package intersection_of_two_arrays

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays

// time: O(n*m) 8ms 77%
// space: S(n+m) 3.1M

// leetcode submit region begin(Prohibit modification and deletion)

func removeDuplicateAndSort(a []int) (ret []int) {
	aLen := len(a)
	sort.Ints(a)
	for i := 0; i < aLen; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func intersection(nums1 []int, nums2 []int) (ret []int) {
	ret = []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	nums1 = removeDuplicateAndSort(nums1)
	nums2 = removeDuplicateAndSort(nums2)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if n1 == n2 {
				ret = append(ret, n1)
			}
		}
	}
	return removeDuplicateAndSort(ret)
}

// leetcode submit region end(Prohibit modification and deletion)
