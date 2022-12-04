package pascals_triangle_ii

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "rowIndex: 0",
			args: args{rowIndex: 0},
			want: []int{1},
		},
		{
			name: "rowIndex: 1",
			args: args{rowIndex: 1},
			want: []int{1, 1},
		},
		{
			name: "rowIndex: 3",
			args: args{rowIndex: 3},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "rowIndex: 4",
			args: args{rowIndex: 4},
			want: []int{1, 4, 6, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
