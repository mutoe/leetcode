package reverse_bits

// https://leetcode.com/problems/reverse-bits

// level: 1
// time: 0ms O(1) 100%
// space: 2.5M O(1) 99.28%

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfLeadingZeros(x uint32) int {
	if x == 0 {
		return 32
	}
	n := 0
	if x <= 0x0000FFFF {
		n = n + 16
		x = x << 16
	}
	if x <= 0x00FFFFFF {
		n = n + 8
		x = x << 8
	}
	if x <= 0x0FFFFFFF {
		n = n + 4
		x = x << 4
	}
	if x <= 0x3FFFFFFF {
		n = n + 2
		x = x << 2
	}
	if x <= 0x7FFFFFFF {
		n = n + 1
	}
	return n
}
func reverseBits(num uint32) uint32 {
	zeroCount := numberOfLeadingZeros(num)
	var result uint32
	for ; num > 0; num /= 2 {
		result <<= 1
		result += num % 2
	}
	return result << zeroCount
}

//leetcode submit region end(Prohibit modification and deletion)
