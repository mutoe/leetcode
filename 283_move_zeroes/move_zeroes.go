package move_zeroes

// https://leetcode.com/problems/move-zeroes

// level: 1
// time: O(n*log(n)) 12ms 6.22%
// space: O(1) 3.8M

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	l := len(nums)
	k := 0
	for i := 0; i < l-1; i++ {
		if k == l {
			return
		}
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[l-1] = 0
			k++
			i--
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
