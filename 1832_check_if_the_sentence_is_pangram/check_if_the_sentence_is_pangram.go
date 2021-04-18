package check_if_the_sentence_is_pangram

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.1M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func checkIfPangram(sentence string) bool {
	var alphabet = [26]bool{}
	for _, char := range sentence {
		alphabet[char-'a'] = true
	}
	for _, b := range alphabet {
		if !b {
			return false
		}
	}
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
