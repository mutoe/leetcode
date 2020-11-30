package baseball_game

import "strconv"

// https://leetcode.com/problems/baseball-game

// level: 1
// time: O(n) 0ms 100%
// space: O(n) 2.6M 31.11%

// leetcode submit region begin(Prohibit modification and deletion)

type Stack struct {
	items []int
}

func NewStack() Stack {
	return Stack{}
}

func (stack *Stack) Push(v int) {
	stack.items = append(stack.items, v)
}

func (stack *Stack) Pop() int {
	if len(stack.items) == 0 {
		return 0
	}
	x := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return x
}
func (stack *Stack) PeekLast(lastN int) int {
	if lastN == 0 {
		lastN = 1
	}
	return stack.items[len(stack.items)-lastN]
}

func calPoints(ops []string) int {
	stack := NewStack()
	for _, op := range ops {
		switch op {
		case "C":
			stack.Pop()
		case "D":
			last := stack.PeekLast(1)
			stack.Push(last * 2)
		case "+":
			last1, last2 := stack.PeekLast(1), stack.PeekLast(2)
			stack.Push(last1 + last2)
		default:
			n, _ := strconv.Atoi(op)
			stack.Push(n)
		}
	}
	sum := 0
	for len(stack.items) > 0 {
		sum += stack.Pop()
	}
	return sum
}

// leetcode submit region end(Prohibit modification and deletion)
