package cont /* Containers */

import (
	"errors"
)

// Enum. Attribute of AVL-tree node, that represents current state of subtrees balance: is one of subtrees higher than another.
type BalanceFactor int

const (
	Balance       BalanceFactor = 0
	LeftIsHigher                = 1
	LeftIsDisbalancedHigher 	= 2
	RightIsHigher               = -LeftIsHigher
	RightIsDisbalancedHigher 	= -LeftIsDisbalancedHigher
)

// TODO docs
type AVLTreeNode[V any] struct {
	value V
	left  *AVLTreeNode[V]
	right *AVLTreeNode[V]
	bf    BalanceFactor
}

// Creates a new AVL-tree node.
func newAVLTreeNode[V any](value V) *AVLTreeNode[V] {
	newNode := &AVLTreeNode[V]{
		value:	value,
		left:	nil,
		right:	nil,
		bf:		Balance
	}

	return newNode
}

// TODO docs
type AVLTree[V any] struct {
	root *AVLTreeNode[V]
	size int
}

// Creates a new empty AVL-tree
func NewAVLTree[V any]() AVLTree[V] {
	return AVLTree[V]{root: nil}
}

// Returns if AVL-tree is empty.
func (t AVLTree[V])IsEmpty() bool {
	return t.root == nil
}

// Returns the number of elements in AVL-tree.
func (t AVLTree[V])Size() int {
	return t.size
}

// TODO docs
func (t AVLTree[V])find(value V) (*AVLTreeNode, bool) {
	var prev *AVLTreeNode[V] = nil
	var cur *AVLTreeNode[V] = t.root
	for cur != nil {
		prev = cur
		if value < cur.value {
			cur = cur.left
		} else if value > cur.value {
			cur = cur.right
		} else {
			return cur, true
		}
	}

	return prev, false
}

// Returns if AVL-tree contains 'value'.
func (t AVLTree[V])Contains(value V) bool {
	_, contains := t.find(value)
	return contains
}

// TODO docs
func (t *AVLTree[V])Insert(value V) {
	parentNode, contains := t.find(value)
	if contains {
		return
	}

	// TODO maybe no need
	if t.IsEmpty() {
		t.root = newAVLTreeNode(value)
		t.size = 1
		return
	}

	// TODO

	t.size++
}

// TODO docs
func (t *AVLTree[V])Delete(value V) {
	valueNode, contains := t.find(value)
	if !contains {
		return
	}

	// TODO

	t.size--
}

// TODO docs
func (t AVLTree[V])Walk(callback func(v V)error) error {
	return nil
}
