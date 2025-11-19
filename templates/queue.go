package templates

// Queue - FIFO (First In First Out)
func queueBasics() {
	queue := []int{}
	// enqueue (add to back)
	queue = append(queue, 1)
	queue = append(queue, 2)
	// dequeue (remove from front)
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		_ = front
	}
	// peek front
	if len(queue) > 0 {
		front := queue[0]
		_ = front
	}
}

// Deque - Double-Ended Queue
type Deque struct {
	items []int
}

func (d *Deque) PushFront(val int) {
	d.items = append([]int{val}, d.items...)
}

func (d *Deque) PushBack(val int) {
	d.items = append(d.items, val)
}

func (d *Deque) PopFront() int {
	val := d.items[0]
	d.items = d.items[1:]
	return val
}

func (d *Deque) PopBack() int {
	val := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return val
}

func (d *Deque) Front() int {
	return d.items[0]
}

func (d *Deque) Back() int {
	return d.items[len(d.items)-1]
}

func (d *Deque) Len() int {
	return len(d.items)
}

// Circular Queue - Fixed Size
type CircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

func NewCircularQueue(k int) *CircularQueue {
	return &CircularQueue{
		data: make([]int, k),
		cap:  k,
	}
}

func (q *CircularQueue) Enqueue(val int) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.rear] = val
	q.rear = (q.rear + 1) % q.cap
	q.size++
	return true
}

func (q *CircularQueue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % q.cap
	q.size--
	return true
}

func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.front]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.size == q.cap
}
