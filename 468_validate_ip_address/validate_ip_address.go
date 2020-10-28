package validate_ip_address

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/validate-ip-address

// level: 2
// time: O(1) 0ms 100%
// space: O(1) 2M 6.67%

// leetcode submit region begin(Prohibit modification and deletion)

func isValidPartOfIpv4(str string) bool {
	if len(str) < 1 || len(str) > 3 {
		return false
	}
	for i, s := range str {
		if i == 0 && len(str) != 1 && s == '0' {
			return false
		}
		if s >= '0' && s <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func isValidPartOfIpv6(str string) bool {
	if len(str) < 1 || len(str) > 4 {
		return false
	}
	for _, s := range str {
		if s >= '0' && s <= '9' {
			continue
		} else if s >= 'A' && s <= 'F' {
			continue
		} else if s >= 'a' && s <= 'f' {
			continue
		}
		return false
	}
	return true
}

func validIPAddress(IP string) string {
	s := strings.Split(IP, ".")
	if len(s) == 4 {
		for _, s2 := range s {
			num, _ := strconv.Atoi(s2)
			if !isValidPartOfIpv4(s2) || num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	} else if len(s) != 1 {
		return "Neither"
	}

	s = strings.Split(IP, ":")
	if len(s) != 8 {
		return "Neither"
	}
	for _, s2 := range s {
		if !isValidPartOfIpv6(s2) {
			return "Neither"
		}
	}
	return "IPv6"
}

// leetcode submit region end(Prohibit modification and deletion)
