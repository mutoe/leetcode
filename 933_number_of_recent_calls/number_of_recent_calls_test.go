package number_of_recent_calls

import (
	"testing"

	"../utils"
)

func TestNumberOfRecentCalls(t *testing.T) {
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
				commands: []string{"RecentCounter", "ping", "ping", "ping", "ping"},
				params:   [][]int{{}, {1}, {100}, {3001}, {3002}},
			},
			want: []int{0, 1, 2, 3, 3},
		},
		{
			name: "test2",
			args: args{
				commands: []string{"RecentCounter", "ping", "ping", "ping", "ping", "ping"},
				params:   [][]int{{}, {642}, {1849}, {4921}, {5936}, {5957}},
			},
			want: []int{0, 1, 2, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]int, len(tt.args.commands))
			var counter RecentCounter
			for i := 0; i < len(tt.args.commands); i++ {
				switch tt.args.commands[i] {
				case "RecentCounter":
					counter = Constructor()
				case "ping":
					ret[i] = counter.Ping(tt.args.params[i][0])
				}
			}
			if !utils.Equal(ret, tt.want) {
				t.Errorf("got %v, but want %v", ret, tt.want)
			}
		})
	}
}
