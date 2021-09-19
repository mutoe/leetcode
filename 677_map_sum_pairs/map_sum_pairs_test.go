package map_sum_pairs

import "testing"

func TestMapSum_Insert(t *testing.T) {
	type args struct {
		key string
		val int
	}
	type Command struct {
		mode int
		args args
		want int
	}
	tests := []struct {
		name     string
		commands []Command
	}{
		{
			name: "test1",
			commands: []Command{
				{1, args{"apple", 3}, 0},
				{2, args{"ap", 0}, 3},
				{1, args{"app", 2}, 0},
				{2, args{"ap", 0}, 5},
			},
		},
		{
			name: "test2",
			commands: []Command{
				{1, args{"apple", 3}, 0},
				{2, args{"ap", 0}, 3},
				{1, args{"app", 2}, 0},
				{1, args{"apple", 2}, 0},
				{2, args{"ap", 0}, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := Constructor()
			for _, c := range tt.commands {
				if c.mode == 1 {
					ms.Insert(c.args.key, c.args.val)
				} else {
					result := ms.Sum(c.args.key)
					if result != c.want {
						t.Errorf("ms.Sum(%v) = %v, want %v", c.args.key, result, c.want)
					}
				}
			}
		})
	}
}
