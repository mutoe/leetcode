package add_strings

import (
	"bytes"
	"strings"
)

// https://leetcode.com/problems/add-strings

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 3M 6.59%

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(s string) string {
	r := []rune(s)
	var output strings.Builder
	for i := len(r) - 1; i >= 0; i-- {
		output.WriteString(string(r[i]))
	}

	return output.String()
}

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	num1, num2 = reverse(num1), reverse(num2)

	zero := uint8('0')
	carry := uint8(0)
	var buffer bytes.Buffer
	for i := 0; i < len(num2); i++ {
		a := num1[i] - zero
		b := num2[i] - zero
		sum := a + b + carry
		carry = 0
		if sum >= 10 {
			carry++
			sum -= 10
		}
		buffer.WriteByte(sum + '0')
	}

	num1Bytes := []byte(num1)
	if carry > 0 {
		for i := len(num2); i < len(num1); i++ {
			if num1[i] == '9' {
				num1Bytes[i] = '0'
				continue
			}

			num1Bytes[i] = num1[i] + 1
			carry = 0
			break
		}
	}
	buffer.Write(num1Bytes[len(num2):])
	if carry > 0 {
		buffer.WriteByte('1')
	}

	return reverse(buffer.String())
}

// leetcode submit region end(Prohibit modification and deletion)
