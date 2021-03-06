package loader

import (
	"testing"
)

// Test loading one word
func TestWordLoading(t *testing.T) {
	root := new(Node)
	loadWord(root, "word", 0)

	if !isWordLoaded(root, "word", 0) {
		t.Errorf("Word \"word\" was not loaded.\n")
	}
}

// Test dictionary loading
func TestDictionaryLoading(t *testing.T) {
	root := LoadDictionary("../../test/test_load.txt")

	// Test loaded words
	words := []string{
		"this",
		"is",
		"a",
		"simple",
		"list",
		"used",
		"to",
		"test",
		"loading",
	}

	for i := 0; i < len(words); i++ {
		if !isWordLoaded(root, words[i], 0) {
			t.Errorf("Word \"%s\" was not loaded.\n", words[i])
		}
	}
}

func isWordLoaded(root *Node, word string, whichChar int) bool {
	if whichChar == len(word) {
		return root.IsWord
	}

	return isWordLoaded(root.Children[word[whichChar]-FIRST_PRINTABLE_ASCII], word, whichChar+1)
}

// Benchmark dictionary loading
func BenchmarkDictionaryLoading(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoadDictionary("../../test/test_words.txt")
	}
}
