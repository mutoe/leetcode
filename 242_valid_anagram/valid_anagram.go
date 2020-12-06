package valid_anagram

// https://leetcode.com/problems/valid-anagram

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 3M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	var m [26]int
	for _, b := range s {
		m[b-'a']++
	}
	for _, b := range t {
		m[b-'a']--
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
