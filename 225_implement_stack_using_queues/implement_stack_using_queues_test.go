package implement_stack_using_queues

import (
	"testing"

	"../utils"
)

func TestImplementStackUsingQueues(t *testing.T) {
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
				commands: []string{"MyStack", "push", "push", "top", "pop", "empty"},
				params:   [][]int{{}, {1}, {2}, {}, {}, {}},
			},
			want: []int{0, 0, 0, 2, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]int, len(tt.args.commands))
			var stack MyStack
			for i := 0; i < len(tt.args.commands); i++ {
				params := tt.args.params[i]
				switch tt.args.commands[i] {
				case "MinStack":
					stack = Constructor()
				case "push":
					for _, param := range params {
						stack.Push(param)
					}
				case "pop":
					ret[i] = stack.Pop()
				case "top":
					ret[i] = stack.Top()
				case "empty":
					if stack.Empty() {
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
