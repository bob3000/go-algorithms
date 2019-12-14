package linkedlist

// List represents the entrypoint to iterate through the nodes
type List struct {
	size     int
	sentinel *node
}

// node represents one element of the linked list
type node struct {
	value interface{}
	next  *node
}

// New returns a new List instance
func New() *List {
	return &List{
		size:     0,
		sentinel: &node{},
	}
}

// Append adds one node with the given value to the end of the list
func (l *List) Append(value interface{}) {
	last := l.sentinel
	for last.next != nil {
		last = last.next
	}
	last.next = &node{
		value,
		nil,
	}
	l.size++
}

// Get returns a value from the list at a given index. The second return value
// says if the index existed
func (l *List) Get(index int) (interface{}, bool) {
	cur := l.sentinel
	for i := -1; i < index; i++ {
		if cur.next == nil {
			return 0, false
		}
		cur = cur.next
	}
	return cur.value, true
}

// Len returns the current list length
func (l *List) Len() int {
	return l.size
}

// Remove takes out the list item at the given position
func (l *List) Remove(index int) (interface{}, bool) {
	if index > l.size-1 {
		return nil, false
	}
	prev := l.sentinel
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retval := prev.next.value
	if prev.next.next != nil {
		*(prev).next = *(prev).next.next
	}
	l.size--
	return retval, true
}

// Insert inserts an item at a given position
func (l *List) Insert(index int, value interface{}) bool {
	if index > l.size {
		return false
	}
	prev := l.sentinel
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	newNode := node{
		value,
		prev.next,
	}
	*(prev).next = newNode
	l.size++
	return true
}
