package sign_of_the_product_of_an_array

// https://leetcode.com/problems/sign-of-the-product-of-an-array

// level: 1
// time: O(n) 4ms 96.21%
// space: O(1) 3.3M 24.24%

//leetcode submit region begin(Prohibit modification and deletion)
func arraySign(nums []int) int {
	result := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			result *= -1
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
