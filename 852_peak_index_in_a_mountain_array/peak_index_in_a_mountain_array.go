package peak_index_in_a_mountain_array

// https://leetcode.com/problems/peak-index-in-a-mountain-array

// level: 1
// time: O(log(n)) 0ms 100%
// space: O(1) 4.8M 14.52%

// leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-1
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			l = mid + 1
		} else if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
			r = mid
		} else {
			return mid
		}
	}
	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
