package structure

// KSmallestTracker efficiently maintains the k smallest elements from a
// dynamic set of integers with O(1) sum queries and O(log n) updates.
//
// It uses two multisets:
//   - smallestK: contains exactly k smallest elements
//   - remaining: contains all other (larger) elements
//
// This allows O(1) access to the sum of k smallest elements while supporting
// efficient insertion and deletion.
type KSmallestTracker struct {
	k         int
	smallestK *MultiSet
	remaining *MultiSet
	sum       int64
}

// NewKSmallestTracker creates a tracker that maintains the k smallest elements.
// The parameter k must be non-negative.
func NewKSmallestTracker(k int) *KSmallestTracker {
	return &KSmallestTracker{
		k:         k,
		smallestK: NewMultiSet(),
		remaining: NewMultiSet(),
		sum:       0,
	}
}

// Add inserts a new element into the tracker.
// Time complexity: O(log n) where n is the total number of unique elements.
func (t *KSmallestTracker) Add(x int) {
	// Determine which set the element belongs to
	if !t.remaining.IsEmpty() {
		if first, ok := t.remaining.First(); ok && x >= first {
			// Element is larger than smallest in remaining set
			t.remaining.Add(x)
		} else {
			// Element belongs to smallestK
			t.smallestK.Add(x)
			t.sum += int64(x)
		}
	} else {
		// No remaining elements, add to smallestK
		t.smallestK.Add(x)
		t.sum += int64(x)
	}
	t.rebalance()
}

// Remove deletes one occurrence of element x from the tracker.
// Time complexity: O(log n) where n is the total number of unique elements.
func (t *KSmallestTracker) Remove(x int) {
	if t.smallestK.Contains(x) {
		t.smallestK.Remove(x)
		t.sum -= int64(x)
	} else if t.remaining.Contains(x) {
		t.remaining.Remove(x)
	}
	t.rebalance()
}

// Sum returns the sum of the k smallest elements currently tracked.
// If fewer than k elements exist, returns the sum of all elements.
// Time complexity: O(1).
func (t *KSmallestTracker) Sum() int64 {
	return t.sum
}

// Size returns the total number of elements (including duplicates) in the tracker.
// Time complexity: O(1).
func (t *KSmallestTracker) Size() int {
	return t.smallestK.Size() + t.remaining.Size()
}

// rebalance ensures smallestK contains exactly k elements (when possible).
// It moves elements between smallestK and remaining to maintain the invariant.
func (t *KSmallestTracker) rebalance() {
	// Move elements from remaining to smallestK if needed
	for t.smallestK.Size() < t.k && !t.remaining.IsEmpty() {
		if x, ok := t.remaining.First(); ok {
			t.remaining.Remove(x)
			t.smallestK.Add(x)
			t.sum += int64(x)
		}
	}
	
	// Move elements from smallestK to remaining if needed
	for t.smallestK.Size() > t.k {
		if x, ok := t.smallestK.Last(); ok {
			t.smallestK.Remove(x)
			t.remaining.Add(x)
			t.sum -= int64(x)
		}
	}
}
