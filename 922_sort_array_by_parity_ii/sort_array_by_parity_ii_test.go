package sort_array_by_parity_ii

import (
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{[]int{4, 2, 5, 7}},
			want: []int{4, 5, 2, 7},
		},
		{
			name: "test2",
			args: args{[]int{2, 3}},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.A); !compareTwoArray(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareTwoArray(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if i%2 == 0 && a[i]%2 == 0 && b[i]%2 == 0 {
			continue
		}
		if i%2 == 1 && a[i]%2 == 1 && b[i]%2 == 1 {
			continue
		}
		return false
	}
	return true
}
