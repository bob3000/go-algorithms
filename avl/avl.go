package avl

import (
	"fmt"
)

type node struct {
	key                 int
	value               interface{}
	isLeftChild         bool
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
		false,
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
			newNode.isLeftChild = false
		} else {
			add(parent.right, newNode)
		}
	} else {
		if parent.left == nil {
			parent.left = newNode
			newNode.parent = parent
			newNode.isLeftChild = true
		} else {
			add(parent.left, newNode)
		}
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
	fmt.Println("rotate left")
	// update right child
	tmp := n.right
	n.right = tmp.left
	tmp.left = n
	if n.right != nil {
		n.right.parent = n
		n.right.isLeftChild = false
	}
	// we are the root node
	if n.parent == nil {
		a.root = tmp
		tmp.parent = nil
		return
	}
	// we are not the root node
	tmp.parent = n.parent
	if n.isLeftChild {
		tmp.isLeftChild = true
		tmp.parent.left = tmp
	} else {
		tmp.isLeftChild = false
		tmp.parent.right = tmp
	}
	n.isLeftChild = true
	n.parent = tmp
}

func (a *Avl) rotateRight(n *node) {
	fmt.Println("rotate right")
	// update left child
	tmp := n.left
	n.left = tmp.right
	tmp.right = n
	if n.left != nil {
		n.left.parent = n
		n.left.isLeftChild = true
	}
	// we are the root node
	if n.parent == nil {
		a.root = tmp
		tmp.parent = nil
		return
	}
	// we are not the root node
	tmp.parent = n.parent
	if n.isLeftChild {
		tmp.isLeftChild = true
		tmp.parent.left = tmp
	} else {
		tmp.isLeftChild = false
		tmp.parent.right = tmp
	}
	tmp.right = n
	n.isLeftChild = false
	n.parent = tmp
}

func (a *Avl) rotate(n *node) {
	if n.isLeftChild {
		if n.parent.isLeftChild {
			// right rotate
			a.rotateRight(n.parent.parent)
		} else {
			// right left rotate
			a.rotateRight(n.parent.parent.right)
			a.rotateLeft(n.parent.parent)
		}
	} else {
		if !n.isLeftChild {
			// left rotate
			a.rotateLeft(n.parent.parent)
		} else {
			// left right rotate
			a.rotateRight(n.parent.parent.left)
			a.rotateLeft(n.parent.parent)
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
