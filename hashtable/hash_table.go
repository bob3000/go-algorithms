package hashtable

import (
	"fmt"

	"github.com/bob3000/structures/linkedlist"
)

// HashTable represents the data structure
type HashTable struct {
	size       int
	numBuckets int
	buckets    []*linkedlist.List
}

func (h *HashTable) bucketOffset(pos int) (bucket, offset int) {
	if pos == 0 {
		return 0, 0
	}
	bucket = pos % h.numBuckets
	offset = pos / h.numBuckets
	return
}

// New returns a new instance of HashTable
func New(numBuckets int) *HashTable {
	return &HashTable{
		0,
		numBuckets,
		make([]*linkedlist.List, numBuckets),
	}
}

// Add puts a value into the belonging bucket
func (h *HashTable) Add(value interface{}) {
	bucket, _ := h.bucketOffset(h.size)
	if h.buckets[bucket] == nil {
		h.buckets[bucket] = linkedlist.New()
	}
	h.buckets[bucket].Append(value)
	h.size++
}

// Get searches and returns an item from the hash table
func (h *HashTable) Get(key int) (interface{}, bool) {
	bucket, offset := h.bucketOffset(key)
	item, ok := h.buckets[bucket].Get(offset)
	if ok {
		return item, true
	}
	return nil, false
}

// Print the data structure to stdout
func (h *HashTable) Print() {
	for i := 0; i < h.size; i++ {
		// fmt.Printf("%d x %d == %t\n", i, h.numBuckets, i == h.numBuckets)
		val, ok := h.Get(i)
		if ok {
			fmt.Print(val)
		} else {
			fmt.Print("err")
		}
		if (i+1)%(h.numBuckets) == 0 {
			fmt.Println()
		} else {
			fmt.Printf("%c", '\t')
		}
	}
	fmt.Println()
}
