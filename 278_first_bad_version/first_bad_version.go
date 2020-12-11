package first_bad_version

// https://leetcode.com/problems/first-bad-version

// level: 1

// time: O(log(n)) 0ms 100%
// space: O(1) 1.9M 100%

var badVersion int

func isBadVersion(version int) bool {
	if version >= badVersion {
		return true
	}
	return false
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// leetcode submit region end(Prohibit modification and deletion)
