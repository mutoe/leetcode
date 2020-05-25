package reverse_vowels_of_a_string

// https://leetcode.com/problems/reverse-vowels-of-a-string

// level: 1
// time: O(n*log(n)) 0ms 100%
// space: O(n) 4M

// leetcode submit region begin(Prohibit modification and deletion)

func isVowel(s byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, vowel := range vowels {
		if vowel == s {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	chars := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isVowel(chars[i]) {
			i++
			continue
		}
		if !isVowel(chars[j]) {
			j--
			continue
		}
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
	return string(chars)
}

// leetcode submit region end(Prohibit modification and deletion)
