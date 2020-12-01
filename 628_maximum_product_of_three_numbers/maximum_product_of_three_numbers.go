package maximum_product_of_three_numbers

import "math"

// https://leetcode.com/problems/maximum-product-of-three-numbers

// level: 1
// time: O(n) 28ms 100%
// space: O(1) 6.3M 11.29%

// leetcode submit region begin(Prohibit modification and deletion)
func insertMax(nums []int, num int) []int {
	if num > nums[0] {
		nums[0], nums[1], nums[2] = num, nums[0], nums[1]
	} else if num > nums[1] {
		nums[1], nums[2] = num, nums[1]
	} else if num > nums[2] {
		nums[2] = num
	}
	return nums
}

func insertMin(nums []int, num int) []int {
	if num < nums[0] {
		nums[1], nums[0] = nums[0], num
	} else if num < nums[1] {
		nums[1] = num
	}
	return nums
}

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	max := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	min := []int{math.MaxInt32, math.MaxInt32}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		max = insertMax(max, num)
		min = insertMin(min, num)
	}

	max1 := max[0] * max[1] * max[2]
	max2 := max[0] * min[0] * min[1]
	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

// leetcode submit region end(Prohibit modification and deletion)
