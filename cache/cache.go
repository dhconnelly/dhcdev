package cache

import (
	"container/heap"
)

type elem struct {
	path  string
	size  int
	count int
}

type frequencyHeap struct {
	elems []elem
}

func (h *frequencyHeap) Len() int {
	return len(h.elems)
}

func (h *frequencyHeap) Less(i, j int) bool {
	return h.elems[i].count < h.elems[j].count
}

func (h *frequencyHeap) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *frequencyHeap) Pop() any {
	last := h.elems[len(h.elems)-1]
	h.elems = h.elems[0 : len(h.elems)-1]
	return last
}

func (h *frequencyHeap) Push(x any) {
	last := x.(elem)
	h.elems = append(h.elems, last)
}

func (h *frequencyHeap) incr(path string) {
	for i := range h.elems {
		if h.elems[i].path == path {
			h.elems[i].count++
			heap.Fix(h, i)
			break
		}
	}
}

type Cache struct {
	data    map[string][]byte
	lfu     *frequencyHeap
	size    int
	maxSize int
}

func (c *Cache) Size() int {
	return c.size
}

func (c *Cache) Len() int {
	return len(c.data)
}

func New(size int) Cache {
	return Cache{
		data:    make(map[string][]byte),
		lfu:     &frequencyHeap{},
		size:    0,
		maxSize: size,
	}
}

func (c *Cache) Get(path string) []byte {
	data, ok := c.data[path]
	if !ok {
		return nil
	}
	c.lfu.incr(path)
	return data
}

func (c *Cache) pop() {
	last := heap.Pop(c.lfu).(elem)
	delete(c.data, last.path)
	c.size -= last.size
}

func (c *Cache) push(path string, data []byte) {
	last := elem{path: path, size: len(data), count: 0}
	heap.Push(c.lfu, last)
	c.data[path] = data
	c.size += last.size
}

func (c *Cache) Put(path string, data []byte) {
	if len(data) > c.maxSize {
		return
	}
	for c.size+len(data) > c.maxSize {
		c.pop()
	}
	c.push(path, data)
}
