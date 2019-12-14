package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	t.Run("add value", func(t *testing.T) {
		size := 10
		h := New(size)
		for i := 1; i < size; i++ {
			h.Add(i, i)
		}
		checkLighter(t, h, 1)
	})
	t.Run("Empty heap", func(t *testing.T) {
		size := 10
		h := New(size)
		if !h.Empty() {
			t.Fatalf("Got false, want true")
		}
		h.Add(1, 1)
		if h.Empty() {
			t.Fatalf("Got true, want false")
		}
	})
	t.Run("Pop value", func(t *testing.T) {
		size := 10
		h := New(size)
		for i := 1; i < size; i++ {
			h.Add(i, i)
		}
		for i := 1; i < size/3; i++ {
			fmt.Print(1)
			h.Pop()
		}
		checkLighter(t, h, 1)
	})
	t.Run("heap sort", func(t *testing.T) {
		numItems := 20
		h := New(numItems)
		for i := 1; i < numItems; i++ {
			h.Add(i, i)
		}
		h.Sort()
		for i := 1; i < numItems; i++ {
			if h.data[i].weight != i {
				t.Fatalf("heap was not properly sorted")
			}
		}
	})
}

func checkLighter(t *testing.T, h *Heap, pos int) {
	t.Helper()
	parentIdx := parent(pos)
	me := h.data[pos]
	if parentIdx > 1 {
		myParent := me.heap.data[parentIdx]
		if me.weight > myParent.weight {
			t.Fatalf("Heap not properly ordered: me (%d) heavier than parent (%d)",
				me.weight, myParent.weight)
		}
	}
	leftChildIdx := left(pos)
	rightChildIdx := right(pos)
	if leftChildIdx <= me.heap.lastPos {
		checkLighter(t, h, leftChildIdx)
	}
	if rightChildIdx <= me.heap.lastPos {
		checkLighter(t, h, rightChildIdx)
	}
}
