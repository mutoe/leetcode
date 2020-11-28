package minimum_subsequence_in_non_increasing_order

// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order

// level: 1
// time: O(n*log(n)) 8ms 35.90%
// space: O(n) 3.2M 84.62%

// leetcode submit region begin(Prohibit modification and deletion)

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, q[l]
	for i < j {
		i++
		j--
		for q[i] > x {
			i++
		}
		for q[j] < x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func Sum(q []int) int {
	sum := 0
	for _, n := range q {
		sum += n
	}
	return sum
}

func minSubsequence(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	quickSort(nums, 0, len(nums)-1)
	sum := Sum(nums)
	i := 0
	for ; i < len(nums); i++ {
		leftSum := Sum(nums[:i+1])
		if leftSum > sum-leftSum {
			break
		}
	}
	return nums[:i+1]
}

// leetcode submit region end(Prohibit modification and deletion)
