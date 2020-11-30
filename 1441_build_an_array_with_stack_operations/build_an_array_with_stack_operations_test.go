package build_an_array_with_stack_operations

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				target: []int{1, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Pop", "Push"},
		},
		{
			name: "test2",
			args: args{
				target: []int{1, 2, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Push"},
		},
		{
			name: "test3",
			args: args{
				target: []int{1, 2},
				n:      4,
			},
			want: []string{"Push", "Push"},
		},
		{
			name: "test4",
			args: args{
				target: []int{2, 3, 4},
				n:      4,
			},
			want: []string{"Push", "Pop", "Push", "Push", "Push"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.target, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
