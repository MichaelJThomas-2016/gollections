package pkg

import (
	"container/heap"
	"sort"
)

// Item represents an element in the heap
type Item[T any] struct {
	Value    T
	Priority float64
	Index    int // Position in the heap
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue[T any] []*Item[T]

// Len Basic heap interface methods
func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// HeapQ implements Python-like heapq functionality
type HeapQ[T any] struct {
	queue PriorityQueue[T]
}

func (h *HeapQ[T]) Len() int {
	return len(h.queue)
}

// NewHeapQ creates a new heap
func NewHeapQ[T any]() *HeapQ[T] {
	return &HeapQ[T]{
		queue: make(PriorityQueue[T], 0),
	}
}

// HeapPush adds an element to the heap
func (h *HeapQ[T]) HeapPush(value T, priority float64) {
	item := &Item[T]{
		Value:    value,
		Priority: priority,
	}
	heap.Push(&h.queue, item)
}

// HeapPop removes and returns the smallest element
func (h *HeapQ[T]) HeapPop() (T, bool) {
	if len(h.queue) == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(&h.queue).(*Item[T])
	return item.Value, true
}

// HeapPushPop pushes a new item and returns the smallest
func (h *HeapQ[T]) HeapPushPop(value T, priority float64) (T, bool) {
	if len(h.queue) == 0 {
		return value, true
	}

	if priority > h.queue[0].Priority {
		old := h.queue[0].Value
		h.queue[0].Value = value
		h.queue[0].Priority = priority
		heap.Fix(&h.queue, 0)
		return old, true
	}
	return value, true
}

// Peek returns the smallest element without removing it
func (h *HeapQ[T]) Peek() (T, bool) {
	if len(h.queue) == 0 {
		var zero T
		return zero, false
	}
	return h.queue[0].Value, true
}

func (h *HeapQ[T]) NLargest(n int, items []T, priorityFn func(T) float64) []T {
	if n <= 0 {
		return []T{}
	}

	// If n is large relative to input size, use sorting
	if n*2 >= len(items) {
		sorted := make([]T, len(items))
		copy(sorted, items)
		sort.Slice(sorted, func(i, j int) bool {
			return priorityFn(sorted[i]) > priorityFn(sorted[j])
		})

		if n > len(sorted) {
			n = len(sorted)
		}
		return sorted[:n]
	}

	// Use min heap to maintain n largest elements
	minHeap := NewHeapQ[T]()

	// Process first n elements
	for i := 0; i < n && i < len(items); i++ {
		priority := priorityFn(items[i])
		minHeap.HeapPush(items[i], priority)
	}

	// Process remaining elements
	smallestPriority := priorityFn(minHeap.queue[0].Value)

	for i := n; i < len(items); i++ {
		currentPriority := priorityFn(items[i])
		if currentPriority > smallestPriority {
			minHeap.HeapPop() // Remove smallest
			minHeap.HeapPush(items[i], currentPriority)
			smallestPriority = priorityFn(minHeap.queue[0].Value)
		}
	}

	// Extract results
	result := make([]T, 0, minHeap.Len())
	for minHeap.Len() > 0 {
		if val, ok := minHeap.HeapPop(); ok {
			result = append(result, val)
		}
	}

	// Reverse to get descending order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// HeapifyFromSlice Helper function to create a heap from slice
func HeapifyFromSlice[T any](items []T, priorityFn func(T) float64) *HeapQ[T] {
	h := NewHeapQ[T]()
	for _, item := range items {
		h.HeapPush(item, priorityFn(item))
	}
	return h
}
