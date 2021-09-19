package design_add_and_search_words_data_structure

import "testing"

func TestWordDictionary_match(t *testing.T) {
	type Command struct {
		mode int
		word string
		want bool
	}
	tests := []struct {
		name     string
		commands []Command
	}{
		{
			name: "test1",
			commands: []Command{
				{1, "bad", false},
				{1, "dad", false},
				{1, "mad", false},
				{2, "pad", false},
				{2, "bad", true},
				{2, ".ad", true},
				{2, "b..", true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd := Constructor()
			for _, c := range tt.commands {
				if c.mode == 1 {
					wd.AddWord(c.word)
				} else {
					result := wd.Search(c.word)
					if result != c.want {
						t.Errorf("wd.Search(%v) = %v, want %v", c.word, result, c.want)
					}
				}
			}

		})
	}
}
