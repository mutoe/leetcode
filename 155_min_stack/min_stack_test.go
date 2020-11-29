package min_stack

import (
	"testing"

	"../utils"
)

func TestMinStack(t *testing.T) {
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
				commands: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
				params:   [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			},
			want: []int{0, 0, 0, 0, -3, 0, 0, -2},
		},
		{
			name: "test2",
			args: args{
				commands: []string{"MinStack", "push", "getMin"},
				params:   [][]int{{}, {-3}, {}},
			},
			want: []int{0, 0, -3},
		},
		{
			name: "test3",
			args: args{
				commands: []string{"MinStack", "push", "push", "push", "top", "pop", "getMin", "pop", "getMin", "pop", "push", "top", "getMin", "push", "top", "getMin", "pop", "getMin"},
				params:   [][]int{{}, {2147483646}, {2147483646}, {2147483647}, {}, {}, {}, {}, {}, {}, {2147483647}, {}, {}, {-2147483648}, {}, {}, {}, {}},
			},
			want: []int{0, 0, 0, 0, 2147483647, 0, 2147483646, 0, 2147483646, 0, 0, 2147483647, 2147483647, 0, -2147483648, -2147483648, 0, 2147483647},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]int, len(tt.args.commands))
			var stack MinStack
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
					stack.Pop()
				case "top":
					ret[i] = stack.Top()
				case "getMin":
					ret[i] = stack.GetMin()
				}
			}
			if !utils.Equal(tt.want, ret) {
				t.Errorf("Got %v, but want %v", ret, tt.want)
			}
		})
	}
}
