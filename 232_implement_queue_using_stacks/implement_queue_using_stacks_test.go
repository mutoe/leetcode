package implement_queue_using_stacks

import (
	"testing"

	"../utils"
)

func TestImplementQueueUsingStacks(t *testing.T) {
	type args struct {
		commands []string
		params   [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				commands: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
				params:   [][]int{{}, {1}, {2}, {}, {}, {}},
			},
			want: []int{0, 0, 0, 1, 1, 0},
		},
		{
			name: "test2",
			args: args{
				commands: []string{"MyQueue", "push", "pop", "empty"},
				params:   [][]int{{}, {1}, {}, {}},
			},
			want: []int{0, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]int, len(tt.args.commands))
			var queue MyQueue
			for i := 0; i < len(tt.args.commands); i++ {
				params := tt.args.params[i]
				switch tt.args.commands[i] {
				case "MyQueue":
					queue = Constructor()
				case "push":
					for _, param := range params {
						queue.Push(param)
					}
				case "pop":
					ret[i] = queue.Pop()
				case "peek":
					ret[i] = queue.Peek()
				case "empty":
					if queue.Empty() {
						ret[i] = 1
					} else {
						ret[i] = 0
					}
				}
			}
			if !utils.Equal(tt.want, ret) {
				t.Errorf("Got %v, but want %v", ret, tt.want)
			}
		})
	}
}
