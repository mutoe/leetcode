package number_of_islands

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name             string
		args             args
		wantIslandsCount int
	}{
		{
			name: "test1",
			args: args{grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			wantIslandsCount: 1,
		},
		{
			name: "test2",
			args: args{grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			wantIslandsCount: 3,
		},
		{
			name:             "test3",
			args:             args{grid: [][]byte{}},
			wantIslandsCount: 0,
		},
		{
			name: "test4",
			args: args{grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			}},
			wantIslandsCount: 1,
		},
		{
			name: "test5",
			args: args{grid: [][]byte{
				{'1', '0', '1'},
				{'1', '1', '1'},
				{'1', '0', '1'},
			}},
			wantIslandsCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIslandsCount := numIslands(tt.args.grid); gotIslandsCount != tt.wantIslandsCount {
				t.Errorf("numIslands() = %v, want %v", gotIslandsCount, tt.wantIslandsCount)
			}
		})
	}
}
