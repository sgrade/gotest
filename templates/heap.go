package templates

import "container/heap"

// Min Heap Implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Max Heap Implementation
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // > for max
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Heap Usage
func heapUsage() {
	// Min Heap
	minH := &MinHeap{}
	heap.Init(minH)
	heap.Push(minH, 3)
	heap.Push(minH, 1)
	heap.Push(minH, 2)
	min := heap.Pop(minH).(int) // 1

	// Max Heap
	maxH := &MaxHeap{}
	heap.Init(maxH)
	heap.Push(maxH, 3)
	heap.Push(maxH, 1)
	heap.Push(maxH, 2)
	max := heap.Pop(maxH).(int) // 3

	_ = min
	_ = max
}

// Top K Pattern
func topK(nums []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

