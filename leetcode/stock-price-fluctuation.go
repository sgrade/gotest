// 2034. Stock Price Fluctuation
// https://leetcode.com/problems/stock-price-fluctuation/

package leetcode

import "container/heap"

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type StockPrice struct {
	tsToPrice        map[int]int
	minHeap          *MinHeap
	maxHeap          *MinHeap
	currentPrice     int
	currentTimestamp int
}

func Constructor() StockPrice {
	tsToPrice := make(map[int]int)
	minHeap := &MinHeap{}
	maxHeap := &MinHeap{}
	currentPrice := 1000000000
	currentTimestamp := 0
	return StockPrice{
		tsToPrice:        tsToPrice,
		minHeap:          minHeap,
		maxHeap:          maxHeap,
		currentPrice:     currentPrice,
		currentTimestamp: currentTimestamp}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	sp.tsToPrice[timestamp] = price
	if timestamp >= sp.currentTimestamp {
		sp.currentPrice = price
		sp.currentTimestamp = timestamp
	}
	heap.Push(sp.minHeap, [2]int{price, timestamp})
	heap.Push(sp.maxHeap, [2]int{-price, timestamp})
}

func (sp *StockPrice) Current() int {
	return sp.currentPrice
}

func (sp *StockPrice) Maximum() int {
	price := (*sp.maxHeap)[0][0]
	timestamp := (*sp.maxHeap)[0][1]
	for sp.tsToPrice[timestamp] != -price {
		heap.Pop(sp.maxHeap)
		price = (*sp.maxHeap)[0][0]
		timestamp = (*sp.maxHeap)[0][1]
	}
	return -price
}

func (sp *StockPrice) Minimum() int {
	price := (*sp.minHeap)[0][0]
	timestamp := (*sp.minHeap)[0][1]
	for sp.tsToPrice[timestamp] != price {
		heap.Pop(sp.minHeap)
		price = (*sp.minHeap)[0][0]
		timestamp = (*sp.minHeap)[0][1]
	}
	return price
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
