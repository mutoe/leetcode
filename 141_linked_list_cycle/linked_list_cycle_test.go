package linked_list_cycle

import (
	"testing"

	. "../utils/list_node"
)

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
			args: args{[]int{1}, -1},
			want: false,
		},
		{
			name: "test4",
			args: args{[]int{1, 3}, -1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ArrayToListNodeWithTailCycle(tt.args.head, tt.args.pos)
			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
