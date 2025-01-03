package gollections

import (
	"reflect"
	"testing"
)

func TestHeapQ_NLargest(t *testing.T) {
	numbers := []int{1, 30, 4, 21, 100, 50, 32, 99, 2, 43}
	heapq := NewHeapQ[int]()

	t.Run("is n largest", func(t *testing.T) {
		expected := []int{100, 99, 50}
		largest := heapq.NLargest(3, numbers, func(x int) float64 { return float64(x) })
		if !reflect.DeepEqual(expected, largest) {
			t.Errorf("Got %v\n", largest)
		}
	})
}
