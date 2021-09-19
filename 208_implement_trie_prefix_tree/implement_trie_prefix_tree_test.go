package implement_trie_prefix_tree

import "testing"

func TestTrie_Search(t *testing.T) {
	type command struct {
		mode string
		arg  string
		want bool
	}

	tests := []struct {
		name     string
		commands []command
	}{
		{
			name: "test1",
			commands: []command{
				{"insert", "apple", false},
				{"search", "apple", true},
				{"search", "app", false},
				{"startsWith", "app", true},
				{"insert", "app", false},
				{"search", "app", true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, c := range tt.commands {
				if c.mode == "insert" {
					trie.Insert(c.arg)
				} else if c.mode == "search" {
					result := trie.Search(c.arg)
					if result != c.want {
						t.Errorf("trie.Search(%v) = %v, but want %v", c.arg, result, c.want)
					}
				} else {
					result := trie.StartsWith(c.arg)
					if result != c.want {
						t.Errorf("trie.StartsWith(%v) = %v, but want %v", c.arg, result, c.want)
					}
				}
			}
		})
	}
}
