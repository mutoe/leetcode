package long_pressed_name

// https://leetcode.com/problems/long-pressed-name

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2M

// leetcode submit region begin(Prohibit modification and deletion)
func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	var i, j int
	var lastChar byte
	for {
		if i == len(name) {
			for ; j < len(typed); j++ {
				if typed[j] != lastChar {
					return false
				}
			}
			return true
		} else if j == len(typed) && i != len(name) {
			return false
		}
		if name[i] != typed[j] {
			if j != len(typed)-1 && typed[j] == lastChar {
				j++
				continue
			}
			return false
		}
		if name[i] == typed[j] {
			lastChar = name[i]
			i++
			j++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
