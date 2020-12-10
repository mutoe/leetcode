package kth_largest_element_in_an_array

// https://leetcode.com/problems/kth-largest-element-in-an-array

// level: 2
// time: O(n*log(n)) 36ms 5.90%
// space: O(1) 4M 35.39%

// leetcode submit region begin(Prohibit modification and deletion)

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l-1, r+1
	for i < j {
		for ok := true; ok; ok = arr[i] < arr[l] {
			i++
		}
		for ok := true; ok; ok = arr[j] > arr[l] {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	quickSort(arr, l, j)
	quickSort(arr, j+1, r)
}

// leetcode submit region end(Prohibit modification and deletion)
