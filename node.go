package bst

// KeyComparator is the interface to be implemented by key comparators.
type KeyComparator interface {
	// Equal returns whether the key i is equal to the key j.
	Equal(i, j interface{}) bool

	// Less returns whether the key i is lower than the key j.
	Less(i, j interface{}) bool
}

// Node represents a Node of a BST.
type Node struct {
	key         interface{}
	val         interface{}
	left, right *Node
}
