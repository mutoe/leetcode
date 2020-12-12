package count_the_number_of_consistent_strings

// https://leetcode.com/problems/count-the-number-of-consistent-strings

// level: 1
// time: O(n^2) 68ms 0%
// space: O(n) 6.8M 0%

// leetcode submit region begin(Prohibit modification and deletion)
func countConsistentStrings(allowed string, words []string) int {
	m := map[int32]bool{}
	for _, char := range allowed {
		m[char] = true
	}
	ret := 0
outer:
	for _, word := range words {
		for _, c := range word {
			if !m[c] {
				continue outer
			}
		}
		ret++
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
