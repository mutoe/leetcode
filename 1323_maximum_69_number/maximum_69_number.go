package maximum_69_number

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/maximum-69-number

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func maximum69Number(num int) int {
	nums := strings.Split(strconv.Itoa(num), "")
	for i := 0; i < len(nums); i++ {
		if nums[i] == "6" {
			nums[i] = "9"
			break
		}
	}
	ret, _ := strconv.Atoi(strings.Join(nums, ""))
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
