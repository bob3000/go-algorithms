package avl

import "testing"

func TestAVL(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		a := New()
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
	t.Run("right rotate", func(t *testing.T) {
		a := New()
		a.Add(10, 10)
		// check rotation with root node involved
		for i := 9; i > 7; i-- {
			a.Add(i, i)
		}
		assert(t, a.root.key, 9)
		assert(t, a.root.left.key, 8)
		assert(t, a.root.right.key, 10)
		// check rotation with lower level parent nodes involved
		for i := 7; i > 5; i-- {
			a.Add(i, i)
		}
		assert(t, a.root.key, 9)
		assert(t, a.root.left.key, 7)
		assert(t, a.root.left.left.key, 6)
		assert(t, a.root.left.right.key, 8)
	})
	t.Run("left rotate", func(t *testing.T) {
		a := New()
		// check rotation with root node involved
		a.Add(10, 10)
		for i := 11; i < 13; i++ {
			a.Add(i, i)
		}
		assert(t, a.root.key, 11)
		assert(t, a.root.left.key, 10)
		assert(t, a.root.right.key, 12)
		// check rotation with lower level parent nodes involved
		for i := 13; i < 16; i++ {
			a.Add(i, i)
		}
		assert(t, a.root.key, 11)
		assert(t, a.root.right.key, 14)
		assert(t, a.root.right.left.key, 13)
		assert(t, a.root.right.right.key, 15)
	})
	t.Run("left right rotate", func(t *testing.T) {
		// a := New()
		// a.Add(20, 20)
		// a.Add(9, 9)
		// a.Add(10, 10)
		// for i := 10; i < 13; i++ {
		// 	a.Add(i, i)
		// }
		// assert(t, a.root.key, 9)
	})
	t.Run("right left rotate", func(t *testing.T) {

	})
}

func assert(t *testing.T, one, two int) {
	t.Helper()
	if one != two {
		t.Fatalf("assert failed %d != %d", one, two)
	}
}
