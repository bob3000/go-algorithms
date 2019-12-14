package heap

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type node struct {
	weight int
	value  interface{}
	heap   *Heap
}

func parent(i int) int {
	return int(i / 2)
}

func left(i int) int {
	return i * 2
}

func right(i int) int {
	return i*2 + 1
}

// Heap is a heap structure
type Heap struct {
	lastPos int
	data    []*node
}

// New returns a new max Heap
func New(initSize int) *Heap {
	return &Heap{
		0,
		make([]*node, initSize),
	}
}

func (h *Heap) rebalance(pos int) {
	parentIdx := parent(pos)
	var parent *node
	if parentIdx < 1 {
		return
	}
	parent = h.data[parentIdx]
	current := h.data[pos]
	if current.weight > parent.weight {
		h.data[pos], h.data[parentIdx] = h.data[parentIdx], h.data[pos]
		// fmt.Printf("%d <=> %d  ", current.weight, parent.weight)
	}
	h.rebalance(parentIdx)
}

// Add another node to the heap
func (h *Heap) Add(weight int, value interface{}) {
	h.lastPos++
	h.data[h.lastPos] = &node{
		weight,
		value,
		h,
	}
	h.rebalance(h.lastPos)
}

// Empty returns true if there are new elements in the heap
func (h *Heap) Empty() bool {
	if h.lastPos == 0 {
		return true
	}
	return false
}

// Pop removes the root element of the heap
func (h *Heap) Pop() (interface{}, error) {
	if h.Empty() {
		return nil, errors.New("Nothing to pop, the heap is empty")
	}
	retval := h.data[1].value
	h.data[1] = h.data[h.lastPos]
	h.lastPos--
	h.trickleDown(1)

	return retval, nil
}

func (h *Heap) trickleDown(pos int) {
	leftIdx := left(pos)
	rightIdx := right(pos)
	var leftChild node
	var rightChild node
	var toSwap int
	if leftIdx < h.lastPos {
		leftChild = *h.data[leftIdx]
	} else {
		leftChild = node{0, 0, h}
	}
	if rightIdx < h.lastPos {
		rightChild = *h.data[rightIdx]
	} else {
		rightChild = node{0, 0, h}
	}
	if rightChild.weight == 0 && leftChild.weight == 0 {
		return
	}
	if rightChild.weight > leftChild.weight {
		toSwap = rightIdx
	} else {
		toSwap = leftIdx
	}
	h.data[pos], h.data[toSwap] = h.data[toSwap], h.data[pos]
	h.trickleDown(toSwap)
}

// Sort returns a sorted copy of the heap
func (h *Heap) Sort() {
	rememberLast := h.lastPos
	for h.lastPos > 0 {
		h.data[1], h.data[h.lastPos] = h.data[h.lastPos], h.data[1]
		h.lastPos--
		h.trickleDown(1)
	}
	h.lastPos = rememberLast
}

// PrintSlice the slice
func (h *Heap) PrintSlice() {
	for i := 1; i <= h.lastPos; i++ {
		fmt.Printf("%d ", h.data[i].weight)
	}
	fmt.Println()
}

// PrintHeap the heap
func (h *Heap) PrintHeap() {
	height := int(math.Ceil(math.Log2(float64(h.lastPos)))) + 1
	width := int(math.Pow(2, float64(height)) * 1)
	itemsCurent := 0
	spaceDiv := 2
	item := 0
outer:
	for i := 0; i < height; i++ {
		itemsCurent = int(math.Pow(2, float64(i)))
		for j := 1; j < itemsCurent; j++ {
			space := int(width / spaceDiv)
			var toPrint int
			if j%2 == 1 {
				item++
				toPrint = h.data[item].weight
				space -= len(strconv.Itoa(toPrint))
			}
			for k := 0; k < space; k++ {
				fmt.Print(" ")
			}
			if j%2 == 1 {
				fmt.Print(toPrint)
			}
			if item >= h.lastPos {
				break outer
			}
		}
		fmt.Printf("\n\n")
		spaceDiv *= 2
	}
	fmt.Println()
}
