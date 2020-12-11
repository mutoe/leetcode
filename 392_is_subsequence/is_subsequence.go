package is_subsequence

// https://leetcode.com/problems/is-subsequence

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M 22.14%

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
