package linked_list_cycle

import "testing"
import . "../utils"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head []int
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{3, 2, 0, -4}, 1},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2}, 0},
			want: true,
		},
		{
			name: "test3",
			args: args{[]int{1}, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ArrayToLinkedListWithTailCycle(tt.args.head, tt.args.pos)
			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*/
