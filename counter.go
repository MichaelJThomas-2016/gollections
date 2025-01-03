package gollections

import (
	"fmt"
	"sort"
)

type Counter[K comparable] struct {
	data map[K]int
}

func NewCounter[K comparable]() *Counter[K] {
	return &Counter[K]{
		data: make(map[K]int),
	}
}

func (c *Counter[K]) Set(key K) {
	_, exists := c.data[key]
	if !exists {
		c.data[key] = 1
	} else {
		c.data[key] += 1
	}
}

func (c *Counter[K]) Get(key K) int {
	return c.data[key]
}

func (c *Counter[K]) Add(key K) {
	c.data[key]++
}

func (c *Counter[K]) AddCount(key K, count int) {
	c.data[key] += count
}

type CountItem[K comparable] struct {
	Key   K
	Count int
}

func (c *Counter[K]) MostCommon(n int) []CountItem[K] {
	items := make([]CountItem[K], 0, len(c.data))
	for k, v := range c.data {
		items = append(items, CountItem[K]{Key: k, Count: v})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return fmt.Sprint(items[i].Key) < fmt.Sprint(items[j].Key)
		}
		return items[i].Count > items[j].Count
	})

	if n > len(items) {
		n = len(items)
	}
	return items[:n]
}
