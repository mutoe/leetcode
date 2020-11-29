package implement_queue_using_stacks

// https://leetcode.com/problems/implement-queue-using-stacks

// level: 1
// time: O(n) 0ms 100%
// space: O(1) 2.1M 6.02%

// leetcode submit region begin(Prohibit modification and deletion)
const N = 100

type MyQueue struct {
	items  [N]int
	start  int
	length int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.items[(this.start+this.length)%N] = x
	this.length++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.length == 0 {
		return 0
	}
	x := this.items[this.start]
	this.start++
	if this.start >= N {
		this.start = 0
	}
	this.length--
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.items[this.start]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// leetcode submit region end(Prohibit modification and deletion)
