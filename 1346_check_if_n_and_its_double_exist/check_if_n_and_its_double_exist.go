package check_if_n_and_its_double_exist

// https://leetcode.com/problems/check-if-n-and-its-double-exist

// level: 1
// time: O(n) 4ms 87.06%
// space: O(n) 3.3M 41.18%

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	zero := false
	for _, num := range arr {
		if num == 0 {
			if zero {
				return true
			} else {
				zero = true
				continue
			}
		}
		if m[num] == true {
			return true
		}
		m[num*2] = true
	}
	for _, num := range arr {
		if m[num] == true {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
