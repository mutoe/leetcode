package find_smallest_letter_greater_than_target

// https://leetcode.com/problems/find-smallest-letter-greater-than-target

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 2.9M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		mid := (l + r) >> 1
		if target >= letters[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[r%len(letters)]
}

// leetcode submit region end(Prohibit modification and deletion)
