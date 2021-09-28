package add_to_array_form_of_integer

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 0, 0}, 34},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{[]int{2, 7, 4}, 181},
			want: []int{4, 5, 5},
		},
		{
			name: "test3",
			args: args{[]int{2, 1, 5}, 806},
			want: []int{1, 0, 2, 1},
		},
		{
			name: "test4",
			args: args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "test5",
			args: args{[]int{1}, 123},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := make([]int, len(tt.args.num))
			for i := range tt.args.num {
				num[i] = tt.args.num[i]
			}
			if got := addToArrayForm(tt.args.num, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm(%v, %v) = %v, want %v", num, tt.args.k, got, tt.want)
			}
		})
	}
}
