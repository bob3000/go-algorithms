package avl

import "testing"

func TestAddFind(t *testing.T) {
	a := New()
	t.Run("add", func(t *testing.T) {
		// root node
		a.Add(5, 5)
		// create level 1 subtrees
		a.Add(3, 3)
		a.Add(7, 7)
		// fill left subtree
		a.Add(2, 2)
		a.Add(4, 4)
		// fill right subtree
		a.Add(6, 6)
		a.Add(8, 8)

		// check root node
		assert(t, a.root.key, 5)
		// check left subtree
		assert(t, a.root.left.key, 3)
		assert(t, a.root.left.left.key, 2)
		assert(t, a.root.left.right.key, 4)

		// check right subtree
		assert(t, a.root.right.key, 7)
		assert(t, a.root.right.left.key, 6)
		assert(t, a.root.right.right.key, 8)
	})
	t.Run("find", func(t *testing.T) {
		_, val := a.Find(5)
		assert(t, val.(int), 5)
		_, val = a.Find(8)
		assert(t, val.(int), 8)
		_, val = a.Find(3)
		assert(t, val.(int), 3)
	})
}

func TestRotateLeft(t *testing.T) {
	a := New()
	t.Run("rotation with root node involved", func(t *testing.T) {
		a.Add(10, 10)
		a.Add(11, 11)
		a.Add(12, 12)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 10)
		assert(t, a.root.right.key, 12)
	})
	t.Run("rotation with lower level parent nodes involved", func(t *testing.T) {
		a.Add(13, 13)
		a.Add(14, 14)
		assert(t, a.root.key, 11)
		assert(t, a.root.right.key, 13)
		assert(t, a.root.right.left.key, 12)
		assert(t, a.root.right.right.key, 14)
	})
}

func TestRotateRight(t *testing.T) {
	a := New()
	t.Run("rotation with root node involved", func(t *testing.T) {
		a.Add(12, 12)
		a.Add(11, 11)
		a.Add(10, 10)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 10)
		assert(t, a.root.right.key, 12)
	})
	t.Run("rotation with lower level parent nodes involved", func(t *testing.T) {
		a.Add(9, 9)
		a.Add(8, 8)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 9)
		assert(t, a.root.left.left.key, 8)
		assert(t, a.root.left.right.key, 10)
	})
}

func TestRotateLeftRight(t *testing.T) {
	a := New()
	t.Run("rotation with root node involved", func(t *testing.T) {
		a.Add(12, 12)
		a.Add(10, 10)
		a.Add(11, 11)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 10)
		assert(t, a.root.right.key, 12)
	})
	t.Run("rotation with lower level parent nodes involved", func(t *testing.T) {
		a.Add(9, 9)
		a.Add(8, 8)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 9)
		assert(t, a.root.left.left.key, 8)
		assert(t, a.root.left.right.key, 10)
	})
}

func TestRotateRightLeft(t *testing.T) {
	a := New()
	t.Run("rotation with root node involved", func(t *testing.T) {
		a.Add(10, 10)
		a.Add(12, 12)
		a.Add(11, 11)
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 10)
		assert(t, a.root.right.key, 12)
	})
	t.Run("rotation with lower level parent nodes involved", func(t *testing.T) {
		a.Add(13, 13)
		a.Add(14, 14)
		assert(t, a.root.key, 11)
		assert(t, a.root.right.key, 13)
		assert(t, a.root.right.left.key, 12)
		assert(t, a.root.right.right.key, 14)
	})
}

func assert(t *testing.T, one, two int) {
	t.Helper()
	if one != two {
		t.Fatalf("assert failed %d != %d", one, two)
	}
}
