package minimum_time_to_type_word_using_special_typewriter

// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const N = 26

func minTimeToType(word string) int {
	ret := 0
	last := 0
	for _, char := range word {
		cur := int(char - 'a')
		diff := abs(cur - last)
		diff = min(N-diff, diff)
		ret += diff + 1
		last = cur
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
