package structure

import "github.com/emirpasic/gods/trees/redblacktree"

// MultiSet is a sorted multiset (bag) data structure that allows duplicate elements.
// It uses a red-black tree for O(log n) sorted access to unique elements and
// a counter map for O(1) count lookups.
type MultiSet struct {
	tree    *redblacktree.Tree
	counter map[int]int
	size    int
}

// NewMultiSet creates and returns a new empty MultiSet.
func NewMultiSet() *MultiSet {
	return &MultiSet{
		tree:    redblacktree.NewWithIntComparator(),
		counter: make(map[int]int),
		size:    0,
	}
}

// Add inserts element x into the multiset.
// Time complexity: O(log n) for first occurrence, O(1) for duplicates.
func (ms *MultiSet) Add(x int) {
	if count, exists := ms.counter[x]; exists {
		ms.counter[x] = count + 1
	} else {
		ms.counter[x] = 1
		ms.tree.Put(x, struct{}{})
	}
	ms.size++
}

// Remove deletes one occurrence of element x from the multiset.
// Returns true if the element was found and removed, false otherwise.
// Time complexity: O(log n) when removing last occurrence, O(1) otherwise.
func (ms *MultiSet) Remove(x int) bool {
	count, exists := ms.counter[x]
	if !exists {
		return false
	}
	
	if count == 1 {
		delete(ms.counter, x)
		ms.tree.Remove(x)
	} else {
		ms.counter[x] = count - 1
	}
	ms.size--
	return true
}

// Size returns the total number of elements in the multiset (including duplicates).
// Time complexity: O(1).
func (ms *MultiSet) Size() int {
	return ms.size
}

// IsEmpty returns true if the multiset contains no elements.
// Time complexity: O(1).
func (ms *MultiSet) IsEmpty() bool {
	return ms.size == 0
}

// First returns the minimum element in the multiset.
// Returns (0, false) if the multiset is empty.
// Time complexity: O(log n).
func (ms *MultiSet) First() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Left().Key.(int), true
}

// Last returns the maximum element in the multiset.
// Returns (0, false) if the multiset is empty.
// Time complexity: O(log n).
func (ms *MultiSet) Last() (int, bool) {
	if ms.tree.Empty() {
		return 0, false
	}
	return ms.tree.Right().Key.(int), true
}

// Contains returns true if element x exists in the multiset.
// Time complexity: O(1).
func (ms *MultiSet) Contains(x int) bool {
	_, exists := ms.counter[x]
	return exists
}

// Count returns the number of occurrences of element x in the multiset.
// Time complexity: O(1).
func (ms *MultiSet) Count(x int) int {
	return ms.counter[x]
}
