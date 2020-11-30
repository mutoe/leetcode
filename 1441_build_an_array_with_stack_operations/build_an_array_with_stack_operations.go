package build_an_array_with_stack_operations

// https://leetcode.com/problems/build-an-array-with-stack-operations

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.3M 100%

// leetcode submit region begin(Prohibit modification and deletion)
func buildArray(target []int, n int) []string {
	ret := make([]string, 0)
	for i, m := 0, 1; i < len(target) && m <= n; m++ {
		ret = append(ret, "Push")
		if target[i] != m {
			ret = append(ret, "Pop")
		} else {
			i++
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
