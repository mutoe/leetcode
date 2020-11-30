package remove_all_adjacent_duplicates_in_string

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string

// level: 1
// time: O(n) 4ms 87.67%
// space: O(n) 5.9M 100%

// leetcode submit region begin(Prohibit modification and deletion)

func removeDuplicates(S string) string {
	stack := make([]int32, len(S))
	length := 0
	for _, char := range S {
		if length > 0 && stack[length-1] == char {
			length--
		} else {
			stack[length] = char
			length++
		}
	}
	return string(stack[:length])
}

// leetcode submit region end(Prohibit modification and deletion)

// Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
// We repeatedly make duplicate removals on S until we no longer can.
// Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
//
// Note:
// 1 <= S.length <= 20000
