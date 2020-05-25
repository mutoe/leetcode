package integer_to_roman

import "strings"

// https://leetcode.com/problems/integer-to-roman

// level: 2
// time: O(log(n)) 12ms 44.82%
// space: O(1) 3.4M

// leetcode submit region begin(Prohibit modification and deletion)
func intToRoman(num int) (ret string) {
	symbolMap := []string{"IV", "XL", "CD", "M0", "00"}
	bit := 0
	for num > 0 {
		tail := num % 10
		one := symbolMap[bit][0]
		five := symbolMap[bit][1]
		ten := symbolMap[bit+1][0]
		switch {
		case tail <= 0:
			break
		case tail <= 3:
			ret = strings.Repeat(string(one), tail) + ret
		case tail == 4:
			ret = string(one) + string(five) + ret
		case tail <= 8:
			ret = string(five) + strings.Repeat(string(one), tail-5) + ret
		case tail == 9:
			ret = string(one) + string(ten) + ret
		}
		num /= 10
		bit++
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)
