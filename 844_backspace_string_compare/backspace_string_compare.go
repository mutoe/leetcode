package backspace_string_compare

// https://leetcode.com/problems/backspace-string-compare

// time: O(n) 0ms 100%
// space: O(1) 2M

// leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	si, sj := 0, 0
	for i >= -1 && j >= -1 {
		if i >= 0 && S[i] == '#' {
			si++
			i--
			continue
		}
		if j >= 0 && T[j] == '#' {
			sj++
			j--
			continue
		}
		if si > 0 {
			if i >= 0 {
				i--
			}
			si--
			continue
		}
		if sj > 0 {
			if j >= 0 {
				j--
			}
			sj--
			continue
		}
		if i >= 0 && j >= 0 && S[i] == T[j] {
			i--
			j--
			continue
		}
		if i == -1 && j == -1 {
			return true
		}
		return false
	}
	if i == -1 && j == -1 {
		return true
	}
	return false
}

// leetcode submit region end(Prohibit modification and deletion)
