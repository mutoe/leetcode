package destination_city

// https://leetcode.com/problems/destination-city

// contest problem
// time: O(n) 0ms 100%
// space: O(n) 4.4M

// leetcode submit region begin(Prohibit modification and deletion)
func destCity(paths [][]string) string {
	pathMap := make(map[string]string)

	for _, path := range paths {
		pathMap[path[0]] = path[1]
	}

	var dest string

	for _, value := range pathMap {
		if dest == "" {
			dest = value
		}
		if val, ok := pathMap[value]; !ok {
			return value
		} else {
			dest = val
		}
	}

	return dest
}

// leetcode submit region end(Prohibit modification and deletion)
