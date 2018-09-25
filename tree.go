package bst

// Tree represents a simple, non self-balancing, Binary Search Tree. It assumes
// all nodes of the tree use the same key type, and are then comparable with the
// provided comparator.
type Tree struct {
	root *Node
	comp KeyComparator
}

// NewTree creates a new Binary Search Tree with a key comparator.
func NewTree(c KeyComparator) *Tree {
	return &Tree{
		comp: c,
	}
}

// Delete is used to remove the node with the provided key from the BST. If the
// node wasn't found it returns false.
func (t *Tree) Delete(key interface{}) bool {
	return t.delete(key)
}

// delete is used to remove a node by its key. 3 situations can take place:
// - the node has no child
// - the node has one child
// - the node has two children
func (t *Tree) delete(key interface{}) bool {
	child, parent := t.search(t.root, nil, key)

	if child == nil {
		return false
	} else if child.left == nil && child.right == nil {
		t.deleteOrphanNode(child, parent)
	} else if child.left == nil || child.right == nil {
		t.deleteNodeWithChild(child, parent)
	} else {
		t.deleteNodeWithChildren(child, parent)
	}

	return true
}

// deleteOrphanNode is used to remove a node with no child.
func (t *Tree) deleteOrphanNode(child, parent *Node) {
	t.replace(child, nil, parent)
}

// deleteWithChild is used to remove a node with one child. The node is
// replaced by its descendant.
func (t *Tree) deleteNodeWithChild(child, parent *Node) {
	if child.left != nil {
		t.replace(child, child.left, parent)
	} else if child.right != nil {
		t.replace(child, child.right, parent)
	}
}

// deleteNodeWithChildren is used to remove a node with two children. First
// we find the in-order successor of the node in its right subtree. Then we copy
// its key and value in place of the node. If this successor has a descendant,
// it's necessarily on its right. In that case we update the refrence on his
// parent to point to its former (right) child.
func (t *Tree) deleteNodeWithChildren(child, parent *Node) {
	successor, sucessorParent := t.findSuccessor(child.right, child)
	child.key = successor.key
	child.val = successor.val

	if successor.right != nil {
		t.replace(successor, successor.right, sucessorParent)
	} else {
		t.replace(successor, nil, sucessorParent)
	}
}

// replace is used to replace an old node by a new node on its parent.
func (t *Tree) replace(old, new, parent *Node) {
	if parent.left == old {
		parent.left = new
	} else if parent.right == old {
		parent.right = new
	}
}

// findSuccessor is used to retrieve the in-order successor of a node.
func (t *Tree) findSuccessor(child, parent *Node) (*Node, *Node) {
	if child.left == nil {
		return child, parent
	}
	return t.findSuccessor(child.left, child)
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

// insert is used to insert a value at the given key within the BST.
func (t *Tree) insert(node *Node, k, v interface{}) {
	if t.comp.Less(k, node.key) {
		t.insertLeft(node, k, v)
	} else {
		t.insertRight(node, k, v)
	}
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

// Search looks for the key within the BST and return the matching node's value
// if found.
func (t *Tree) Search(key interface{}) (interface{}, bool) {
	child, _ := t.search(t.root, nil, key)
	if child == nil {
		return nil, false
	}
	return child.val, true
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
