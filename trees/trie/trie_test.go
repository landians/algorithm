package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func initTrie() *Trie {
	tree := NewTrie()

	tree.Insert("hello")
	tree.Insert("heiooio")
	tree.Insert("just")
	tree.Insert("justinbieber")
	return tree
}

func TestTrie_PrefixNumber(t *testing.T) {
	tree := initTrie()
	assert.Equal(t, tree.PrefixNumber("he"), 2)
	assert.Equal(t, tree.PrefixNumber("be"), 0)
}

func TestTrie_Search(t *testing.T) {
	tree := initTrie()
	assert.Equal(t, tree.Search("hello"), 1)
	assert.Equal(t, tree.Search("he"), 0)
}

func TestTrie_Delete(t *testing.T) {
	tree := initTrie()
	tree.Delete("hello")
	assert.Equal(t, tree.Search("hello"), 0)
}
