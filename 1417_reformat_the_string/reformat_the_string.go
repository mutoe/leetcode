package reformat_the_string

// https://leetcode.com/problems/reformat-the-string

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 3.4M 73.68%

//leetcode submit region begin(Prohibit modification and deletion)
import "strings"

func reformat(s string) string {
	letters := strings.Builder{}
	digits := strings.Builder{}
	for _, char := range s {
		if char >= '0' && char <= '9' {
			digits.WriteRune(char)
		} else {
			letters.WriteRune(char)
		}
	}
	diff := letters.Len() - digits.Len()
	if diff > 1 || diff < -1 {
		return ""
	}
	if diff > 0 {
		return buildString(letters, digits)
	} else if diff < 0 {
		return buildString(digits, letters)
	}
	if s[0] >= '0' && s[0] <= '9' {
		return buildString(letters, digits)
	} else {
		return buildString(digits, letters)
	}
}

func buildString(sb1, sb2 strings.Builder) string {
	ret := strings.Builder{}
	s1, s2 := sb1.String(), sb2.String()
	for i := 0; i < sb1.Len(); i++ {
		ret.WriteByte(s1[i])
		if i < sb2.Len() {
			ret.WriteByte(s2[i])
		}
	}
	return ret.String()
}

//leetcode submit region end(Prohibit modification and deletion)
