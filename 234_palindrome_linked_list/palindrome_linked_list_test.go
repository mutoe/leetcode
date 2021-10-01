package palindrome_linked_list

import (
	"testing"

	. "../utils/list_node"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{1, 2}},
			want: false,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 2, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(ArrayToListNode(tt.args.head)); got != tt.want {
				//t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.head, got, tt.want)
			}
		})
	}
}
