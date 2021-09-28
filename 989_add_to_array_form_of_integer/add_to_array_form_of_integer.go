package add_to_array_form_of_integer

// https://leetcode.com/problems/add-to-array-form-of-integer

// level: 1
// time: O(n) 24ms 100%
// space: O(1) 6.7M 83.33%

//leetcode submit region begin(Prohibit modification and deletion)
func addToArrayForm(num []int, k int) []int {
	i := len(num) - 1
	for k > 0 {
		cur := k % 10
		k /= 10
		if num[i]+cur >= 10 {
			k += 1
		}
		num[i] = (num[i] + cur) % 10
		if i > 0 {
			i--
		} else if k != 0 {
			num = append([]int{0}, num...)
		}
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
