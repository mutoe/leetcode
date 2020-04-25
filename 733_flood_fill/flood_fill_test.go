package flood_fill

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{
			name: "test1",
			args: args{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			wantResult: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "test2",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 1, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 1,
			},
			wantResult: [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("floodFill() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
