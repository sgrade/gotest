// 703. Kth Largest Element in a Stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream/

package leetcode

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k       int
	minHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := &KthLargest{
		k:       k,
		minHeap: &MinHeap{},
	}
	for _, num := range nums {
		heap.Push(kthLargest.minHeap, num)
		if kthLargest.minHeap.Len() > k {
			heap.Pop(kthLargest.minHeap)
		}
	}
	return *kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.minHeap.Len() < this.k {
		heap.Push(this.minHeap, val)
	} else if val > (*this.minHeap)[0] {
		heap.Push(this.minHeap, val)
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
