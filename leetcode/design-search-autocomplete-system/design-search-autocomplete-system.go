// 642. Design Search Autocomplete System
// https://leetcode.com/problems/design-search-autocomplete-system/

package designsearchautocompletesystem

import "sort"

// Based on Editorial's Approach 1: Trie
// Each trie node stores children and a frequency map of every sentence
// that passes through it, enabling O(1) candidate lookup at any prefix.
type trieNode struct {
	children  map[byte]*trieNode
	sentences map[string]int
}

func newTrieNode() *trieNode {
	return &trieNode{
		children:  make(map[byte]*trieNode),
		sentences: make(map[string]int),
	}
}

type AutocompleteSystem struct {
	root   *trieNode
	curr   *trieNode
	dead   *trieNode // sink node for prefixes not in the trie
	prefix []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	root := newTrieNode()
	sys := AutocompleteSystem{
		root: root,
		curr: root,
		dead: newTrieNode(),
	}
	for i, s := range sentences {
		sys.addToTrie(s, times[i])
	}
	return sys
}

// addToTrie inserts sentence and propagates its count to every node along the path.
func (s *AutocompleteSystem) addToTrie(sentence string, count int) {
	node := s.root
	for i := range sentence {
		c := sentence[i]
		if _, ok := node.children[c]; !ok {
			node.children[c] = newTrieNode()
		}
		node = node.children[c]
		node.sentences[sentence] += count
	}
}

func (s *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		s.addToTrie(string(s.prefix), 1)
		s.prefix = s.prefix[:0]
		s.curr = s.root
		return nil
	}

	s.prefix = append(s.prefix, c)
	next, ok := s.curr.children[c]
	if !ok {
		s.curr = s.dead
		return nil
	}
	s.curr = next

	// Sort by frequency descending, then lexicographically, and return top 3.
	cands := make([]string, 0, len(s.curr.sentences))
	for sent := range s.curr.sentences {
		cands = append(cands, sent)
	}
	sort.Slice(cands, func(i, j int) bool {
		if fi, fj := s.curr.sentences[cands[i]], s.curr.sentences[cands[j]]; fi != fj {
			return fi > fj
		}
		return cands[i] < cands[j]
	})
	if len(cands) > 3 {
		cands = cands[:3]
	}
	return cands
}
