package bst

// Comparator is the interface to be implemented by key comparators.
type Comparator interface {
	Less(i, j interface{}) bool
	Equal(i, j interface{}) bool
}

// Node represents a Node of a BST.
type Node struct {
	key   interface{}
	val   interface{}
	left  *Node
	right *Node
}
