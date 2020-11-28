package increasing_decreasing_string

import "sort"

// https://leetcode.com/problems/increasing-decreasing-string

// level: 1
// time: O(n*log(n)) 12ms 30.77%
// space: O(n) 3.3M 64.10%

// leetcode submit region begin(Prohibit modification and deletion)
func sortString(s string) string {
	bytes := []byte(s)
	ret := make([]byte, 0)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	for len(ret) < len(s) {
		i := 0
		var lastChar byte
		for i < len(bytes) {
			if lastChar != bytes[i] {
				ret = append(ret, bytes[i])
				lastChar = bytes[i]
				bytes = append(bytes[:i], bytes[i+1:]...)
			} else {
				i++
			}
		}
		i--
		lastChar = 0
		for i >= 0 {
			if lastChar != bytes[i] {
				ret = append(ret, bytes[i])
				lastChar = bytes[i]
				bytes = append(bytes[:i], bytes[i+1:]...)
			}
			i--
		}
	}

	return string(ret)
}

// leetcode submit region end(Prohibit modification and deletion)
