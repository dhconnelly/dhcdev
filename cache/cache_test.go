package cache

import (
	"slices"
	"testing"
)

func constant(size int, value byte) []byte {
	b := make([]byte, size)
	for i := range b {
		b[i] = value
	}
	return b
}

func TestCache(t *testing.T) {
	maxSize := 2048
	cache := New(maxSize)

	data1 := constant(1024, 1)
	data2 := constant(1024, 2)
	data3 := constant(1024, 3)

	// cache should be empty
	if cache.size != 0 || len(cache.data) != 0 || cache.Len() != 0 {
		t.Fatalf("cache should be empty!")
	}

	// add two items, they should fit
	cache.Put("data1", data1)
	cache.Put("data2", data2)
	if cache.size != len(data1)+len(data2) ||
		!slices.Equal(cache.Get("data1"), data1) ||
		!slices.Equal(cache.Get("data2"), data2) {
		t.Fatalf("cache should contain added elements")
	}

	// ping both of them a couple of times
	cache.Get("data1") // 2
	cache.Get("data1") // 3
	cache.Get("data1") // 4
	cache.Get("data2") // 1
	cache.Get("data2") // 2

	// add a new element, should evict the previous with lower count
	cache.Put("data3", data3)
	if cache.size != len(data1)+len(data3) ||
		!slices.Equal(cache.Get("data1"), data1) ||
		!slices.Equal(cache.Get("data3"), data3) ||
		cache.Get("data2") != nil {
		t.Fatalf("cache should have evicted data2")
	}

	// add the evicted element, should evict the newer element
	cache.Put("data2", data2)
	if cache.size != len(data1)+len(data2) ||
		!slices.Equal(cache.Get("data1"), data1) ||
		!slices.Equal(cache.Get("data2"), data2) ||
		cache.Get("data3") != nil {
		t.Fatalf("cache should have evicted data3")
	}

	// ping the lower count element until it's higher
	cache.Get("data2") // 2
	cache.Get("data2") // 3
	cache.Get("data2") // 4
	cache.Get("data2") // 5
	cache.Get("data2") // 6
	cache.Get("data2") // 7

	// add again, should evict first element
	cache.Put("data3", data3)
	if cache.size != len(data2)+len(data3) ||
		!slices.Equal(cache.Get("data2"), data2) ||
		!slices.Equal(cache.Get("data3"), data3) ||
		cache.Get("data1") != nil {
		t.Fatalf("cache should have evicted data1")
	}
}
