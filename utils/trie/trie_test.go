package trie

import "testing"

func TestTrie_contains(t *testing.T) {

	t.Run("add", func(t *testing.T) {
		trie := &Trie{}

		trie.Add("mutoe")
		if trie.size != 1 {
			t.Errorf("trie size should equal %v, but got %v", 1, trie.size)
		}

		if result := trie.Contains("mutoe"); !result {
			t.Errorf("should have mutoe but not")
		}
	})
}
