package implement_stack_using_queues

// https://leetcode.com/problems/implement-stack-using-queues

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2M 100%

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	items []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.items[len(this.items)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.items) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
