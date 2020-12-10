package kth_largest_element_in_an_array

// https://leetcode.com/problems/kth-largest-element-in-an-array

// level: 2
// time: O(n*log(n)) 4ms 98.66%
// space: O(1) 3.5M 58.98%

// TODO: Need to review

// leetcode submit region begin(Prohibit modification and deletion)

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(arr []int, l, r, k int) int {
	mid := (l + r) >> 1
	arr[l], arr[mid] = arr[mid], arr[l]

	p := pivot(arr, l, r)
	if k == p {
		return arr[p]
	}
	if k < p {
		return quickSort(arr, l, p-1, k)
	} else {
		return quickSort(arr, p+1, r, k)
	}
}

func pivot(arr []int, l, r int) int {
	// arr[l+1: i] <= v;  arr[j+1: r+1] >= v
	i, j := l+1, r
	for {
		for i <= j && arr[i] < arr[l] {
			i++
		}
		for i <= j && arr[j] > arr[l] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i, j = i+1, j-1
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// leetcode submit region end(Prohibit modification and deletion)
