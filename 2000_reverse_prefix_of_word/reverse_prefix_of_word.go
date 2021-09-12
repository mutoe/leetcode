package reverse_prefix_of_word

// https://leetcode.com/problems/reverse-prefix-of-word

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.1M 100%

//leetcode submit region begin(Prohibit modification and deletion)
func reversePrefix(word string, ch byte) string {
	i := 0
	for ; i < len(word); i++ {
		if word[i] == ch {
			break
		}
	}
	if i == len(word) {
		return word
	}
	str := []byte(word)
	for j := 0; j < i; j, i = j+1, i-1 {
		str[i], str[j] = word[j], word[i]
	}
	return string(str)
}

//leetcode submit region end(Prohibit modification and deletion)
