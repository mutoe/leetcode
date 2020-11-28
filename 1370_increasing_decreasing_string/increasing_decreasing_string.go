package increasing_decreasing_string

// https://leetcode.com/problems/increasing-decreasing-string

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 3.4M 56.41%

// leetcode submit region begin(Prohibit modification and deletion)
func sortString(s string) string {
	var chars [26]byte
	ret := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	for len(ret) < len(s) {
		i := 0
		for ; i < 26; i++ {
			if chars[i] > 0 {
				ret = append(ret, byte(i)+'a')
				chars[i]--
			}
		}
		i--
		for ; i >= 0; i-- {
			if chars[i] > 0 {
				ret = append(ret, byte(i)+'a')
				chars[i]--
			}
		}
	}

	return string(ret)
}

// leetcode submit region end(Prohibit modification and deletion)
