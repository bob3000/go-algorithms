package linkedlist

import "testing"

func TestList(t *testing.T) {
	t.Run("append item", func(t *testing.T) {
		l := New()
		l.Append(1)
		got, ok := l.Get(1)
		if got != 1 && ok {
			t.Fatalf("Appended value should be 1, got %d", got)
		}
	})
	t.Run("length", func(t *testing.T) {
		l := New()
		got := l.Len()
		if got != 0 {
			t.Fatalf("Length should be 0, got %d", got)
		}
		l.Append(1)
		got = l.Len()
		if got != 1 {
			t.Fatalf("Length should be 1, got %d", got)
		}
	})
	t.Run("remove item", func(t *testing.T) {
		l := New()
		l.Append(0)
		l.Append(1)
		got, ok := l.Get(0)
		if got != 0 && ok {
			t.Fatalf("Precondition not met, expected element 0 to be 0, "+
				"got %d and ok %t", got, ok)
		}
		l.Remove(0)
		got, _ = l.Get(0)
		if got != 1 {
			t.Fatalf("Element 0 should be 1, got %d", got)
		}
	})
	t.Run("insert item", func(t *testing.T) {
		l := New()
		l.Append(1)
		got, ok := l.Get(0)
		if got != 1 && ok {
			t.Fatalf("Precondition not met, expected element 0 to be 0, "+
				"got %d and ok %t", got, ok)
		}
		l.Insert(0, 0)
		got, _ = l.Get(0)
		if got != 0 {
			t.Fatalf("Element 0 should be 0, got %d", got)
		}
	})
}
