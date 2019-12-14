package hashtable

import "testing"

func TestHashMap(t *testing.T) {
	t.Run("add item", func(t *testing.T) {
		h := New(3)
		for i := 0; i < 6; i++ {
			h.Add(i)
		}
		for i := 0; i < 6; i++ {
			got, ok := h.Get(i)
			want := i
			if !ok || got != want {
				t.Fatalf("Got wrong number, got %d, want %d", got, want)
			}
		}
	})
}
