package avl

type node struct {
	key                 int
	value               interface{}
	parent, left, right *node
}

// Avl is the basic type of the AVL tree
type Avl struct {
	root *node
	size int
}

// New creates a new AVL tree
func New() Avl {
	return Avl{}
}

// Height recursively measures the hight of the tree
func (a *Avl) Height() int {
	return height(a.root)
}

func height(n *node) int {
	if n == nil {
		return 0
	}
	heightLeft := height(n.left) + 1
	heightRight := height(n.right) + 1
	if heightLeft > heightRight {
		return heightLeft
	}
	return heightRight
}

// Add aggregates another node to the tree
func (a *Avl) Add(key int, value interface{}) {
	n := &node{
		key,
		value,
		nil, nil, nil,
	}
	a.size++

	if a.root == nil {
		a.root = n
		return
	}
	add(a.root, n)
	if !a.isBalanced() {
		a.rotate(n)
	}
}

func add(parent, newNode *node) {
	if parent.key < newNode.key {
		if parent.right == nil {
			parent.right = newNode
			newNode.parent = parent
		} else {
			add(parent.right, newNode)
		}
	} else {
		if parent.left == nil {
			parent.left = newNode
			newNode.parent = parent
		} else {
			add(parent.left, newNode)
		}
	}
}

// Remove the node with the given key and return it's value along with
// a boolean which indicates if the removal was successful
func (a *Avl) Remove(key int) (bool, interface{}) {
	ok, n := remove(a.root, key)
	if !ok {
		return false, nil
	}
	return true, n.value
}

func remove(n *node, key int) (bool, *node) {
	ok, n := find(n, key)
	if !ok {
		return ok, nil
	}

	var parentPointer **node
	if isLeftChild := n.parent.left == n; isLeftChild {
		parentPointer = &n.parent.left
	} else {
		parentPointer = &n.parent.right
	}
	if n.left == nil && n.right == nil {
		*parentPointer = nil
	} else if n.left == nil && n.right != nil {
		*parentPointer = n.right
		n.right.parent = *parentPointer
	} else if n.left != nil && n.right == nil {
		*parentPointer = n.left
		n.left.parent = *parentPointer
	} else {
		_, m := min(n.right)
		n.key, n.value = m.key, m.value
		remove(m, key)
	}
	n.parent = nil
	return true, n
}

// Min returns the tree's minimum value
func (a *Avl) Min() (bool, interface{}) {
	ok, n := min(a.root)
	if !ok {
		return false, nil
	}
	return true, n.value
}

func min(n *node) (bool, *node) {
	if n == nil {
		return false, nil
	} else if n.left == nil {
		return true, n
	} else {
		return min(n.left)
	}
}

// Max returns the tree's minimum value
func (a *Avl) Max() (bool, interface{}) {
	ok, n := max(a.root)
	if !ok {
		return false, nil
	}
	return true, n.value
}

func max(n *node) (bool, *node) {
	if n == nil {
		return false, nil
	} else if n.right == nil {
		return true, n
	} else {
		return max(n.right)
	}
}

// Find returns the value of a given a key and a boolean value
// indicating if the key was even found
func (a *Avl) Find(key int) (bool, interface{}) {
	ok, n := find(a.root, key)
	if !ok {
		return false, nil
	}
	return true, n.value
}

func find(n *node, keyToFind int) (bool, *node) {
	if n == nil {
		return false, nil
	} else if n.key == keyToFind {
		return true, n
	} else if n.key < keyToFind {
		return find(n.right, keyToFind)
	} else {
		return find(n.left, keyToFind)
	}
}

func (a *Avl) isBalanced() bool {
	heightLeft := height(a.root.left)
	heightRight := height(a.root.right)
	if heightLeft-heightRight > 1 || heightRight-heightLeft > 1 {
		return false
	}
	return true
}

func (a *Avl) rotateLeft(n *node) {
	tmp := n.right
	// tmp left subtree
	n.right = tmp.left
	if tmp.left != nil {
		tmp.left.parent = n
	}
	// tmp
	tmp.parent = n.parent
	if tmp.parent == nil {
		a.root = tmp
	} else if tmp.parent.left == n {
		tmp.parent.left = tmp
	} else {
		tmp.parent.right = tmp
	}
	tmp.left = n
	n.parent = tmp
}

func (a *Avl) rotateRight(n *node) {
	tmp := n.left
	// tmp left subtree
	n.left = tmp.right
	if tmp.right != nil {
		tmp.right.parent = n
	}
	// tmp
	tmp.parent = n.parent
	if tmp.parent == nil {
		a.root = tmp
	} else if tmp.parent.right == n {
		tmp.parent.right = tmp
	} else {
		tmp.parent.left = tmp
	}
	tmp.right = n
	n.parent = tmp
}

func (a *Avl) rotate(n *node) {
	if n.parent.left == n {
		if n.parent.parent.left != nil &&
			n.parent.parent.left.left == n {
			// right rotate
			a.rotateRight(n.parent.parent)
		} else {
			// right left rotate
			a.rotateRight(n.parent.parent.right)
			a.rotateLeft(n.parent)
		}
	} else {
		if n.parent.parent.right != nil &&
			n.parent.parent.right.right == n {
			// left rotate
			a.rotateLeft(n.parent.parent)
		} else {
			// left right rotate
			a.rotateLeft(n.parent.parent.left)
			a.rotateRight(n.parent)
		}
	}
}

// PreOrderTraversal traverses in pre order
func PreOrderTraversal(n *node, f func(n interface{})) {
	if n == nil {
		return
	}
	f(n.value)
	PreOrderTraversal(n.left, f)
	PreOrderTraversal(n.right, f)
}

// InOrderTraversal traverses in order
func InOrderTraversal(n *node, f func(n interface{})) {
	if n == nil {
		return
	}
	InOrderTraversal(n.left, f)
	f(n.value)
	InOrderTraversal(n.right, f)
}

// PostOrderTraversal traverses in post order
func PostOrderTraversal(n *node, f func(n interface{})) {
	if n == nil {
		return
	}
	PostOrderTraversal(n.left, f)
	PostOrderTraversal(n.right, f)
	f(n.value)
}
