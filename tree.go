package bst

// Tree represents a simple, non self-balancing, Binary Search Tree. It assumes
// all nodes of the tree use the same key type, and are then comparable with the
// provided comparator.
type Tree struct {
	root *Node
	comp Comparator
}

// NewTree creates a new Binary Search Tree with c as the key Comparator.
func NewTree(c Comparator) *Tree {
	return &Tree{
		comp: c,
	}
}

// Insert is used to add a new node to the BST or replace the value of an
// existing node sharing the same key.
func (t *Tree) Insert(key, val interface{}) {
	if t.root == nil {
		t.root = &Node{
			key: key,
			val: val,
		}
		return
	}
	t.insert(t.root, key, val)
}

// findSuccessor is used to retrieve the in-order successor of a node.
func (t *Tree) findSuccessor(child, parent *Node) (*Node, *Node) {
	if child.left == nil {
		return child, parent
	}
	return t.findSuccessor(child.left, child)
}

// Search looks for the key within the BST and return the matching node's value
// if found.
func (t *Tree) Search(key interface{}) (interface{}, bool) {
	child, _ := t.search(t.root, nil, key)
	if child == nil {
		return nil, false
	}
	return child.val, true
}

// insert is used to insert a value at the given key within the BST.
func (t *Tree) insert(node *Node, k, v interface{}) {
	if t.comp.Less(k, node.key) {
		t.insertLeft(node, k, v)
		return
	}
	t.insertRight(node, k, v)
}

// insertLeft is used to insert a value into the left branch of a subtree.
func (t *Tree) insertLeft(parent *Node, k, v interface{}) {
	if parent.left == nil {
		parent.left = &Node{
			key: k,
			val: v,
		}
		return
	}
	t.insert(parent.left, k, v)
}

// insertRight is used to insert a value into the right branch of a subtree.
func (t *Tree) insertRight(parent *Node, k, v interface{}) {
	if parent.right == nil {
		parent.right = &Node{
			key: k,
			val: v,
		}
		return
	}
	t.insert(parent.right, k, v)
}

// search is used to lookup a key in the BST and return the matching node and
// its parent.
func (t *Tree) search(child, parent *Node, k interface{}) (*Node, *Node) {
	if child == nil || t.comp.Equal(k, child.key) {
		return child, parent
	}
	if t.comp.Less(k, child.key) {
		return t.search(child.left, child, k)
	}
	return t.search(child.right, child, k)
}
