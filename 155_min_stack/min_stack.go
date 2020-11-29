package min_stack

import "math"

// https://leetcode.com/problems/min-stack

// level: 1
// time: O(n) 12ms 98.73%
// space: O(n) 7.5M 50.63%

// leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	items []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := MinStack{
		min: math.MaxInt32,
	}
	return s
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}
	lastIndex := len(this.items) - 1
	x := this.items[lastIndex]
	this.items = this.items[:lastIndex]
	if len(this.items) == 0 {
		this.min = math.MaxInt32
	} else if x == this.min {
		min := this.items[0]
		for _, value := range this.items {
			if value < min {
				min = value
			}
		}
		this.min = min
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// leetcode submit region end(Prohibit modification and deletion)
