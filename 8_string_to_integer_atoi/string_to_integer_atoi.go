package string_to_integer_atoi

import "math"

// https://leetcode.com/problems/string-to-integer-atoi

// level: 2
// time: O(n) 4ms 47.51%
// space: O(1) 2.3M

// leetcode submit region begin(Prohibit modification and deletion)
func isSupportedChars(char int32) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func myAtoi(str string) (ret int) {
	sign := 0
	for _, char := range str {
		if sign == 0 && char == ' ' {
			continue
		}
		if sign == 0 && char == '0' {
			sign = 1
			continue
		}
		if sign == 0 && char == '+' {
			sign = 1
			continue
		}
		if sign == 0 && char == '-' {
			sign = -1
			continue
		}
		if isSupportedChars(char) {
			if sign == 0 {
				sign = 1
			}
			ret = ret*10 + int(char-'0')
			if ret > math.MaxInt32 {
				break
			}
			continue
		}
		break
	}
	if sign == 0 {
		sign = 1
	}
	ret = ret * sign
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
