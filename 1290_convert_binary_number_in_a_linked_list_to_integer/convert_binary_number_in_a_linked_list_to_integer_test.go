package convert_binary_number_in_a_linked_list_to_integer

import "testing"
import . "../utils/list_node"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]int{1, 0, 1}},
			want: 5,
		},
		{
			name: "test2",
			args: args{[]int{0}},
			want: 0,
		},
		{
			name: "test3",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "test4",
			args: args{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}},
			want: 18880,
		},
		{
			name: "test5",
			args: args{[]int{0, 0}},
			want: 0,
		},
		{
			name: "test6",
			args: args{[]int{1, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(ArrayToListNode(tt.args.head)); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
