package check_if_n_and_its_double_exist

import "testing"

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{[]int{10, 2, 5, 3}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{7, 1, 14, 11}},
			want: true,
		},
		{
			name: "test3",
			args: args{[]int{3, 1, 7, 11}},
			want: false,
		},
		{
			name: "test4",
			args: args{[]int{-2, 0, 10, -19, 4, 6, -8}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
