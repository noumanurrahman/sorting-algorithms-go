package structures

import "fmt"

type RBNode struct {
	red    bool
	parent *RBNode
	left   *RBNode
	right  *RBNode
	value  *int
}

func (node *RBNode) New(value int) {
	node.value = &value
}

type RBTree struct {
	root *RBNode
	null *RBNode
}

func (tree *RBTree) New() {
	node := RBNode{value: nil}
	tree.null = &node
	tree.root = tree.null
}

func (tree *RBTree) Insert(value int) {
	newNode := RBNode{value: &value, red: true}
	current := tree.root
	var parent *RBNode
	for current != nil && current != tree.null {
		parent = current
		if *newNode.value < *current.value {
			current = current.left
		} else if *newNode.value > *current.value {
			current = current.right
		} else {
			return
		}
	}
	newNode.parent = parent
	if parent == nil {
		tree.root = &newNode
	} else {
		if *parent.value < *newNode.value {
			parent.right = &newNode
		} else if *parent.value > *newNode.value {
			parent.left = &newNode
		}
	}
	tree.FixInsert(&newNode)
}

func (tree *RBTree) FixInsert(node *RBNode) {
	current := node
	for tree.root != current && current.parent.red {
		parent := current.parent
		grandparent := parent.parent
		if parent == grandparent.right {
			uncle := grandparent.left
			if uncle != nil && uncle.red {
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent
			} else {
				if current == parent.left {
					current = parent
					tree.RotateRight(current)
					parent = current.parent
				}
				parent.red = false
				grandparent.red = true
				tree.RotateLeft(grandparent)
			}
		} else if parent == grandparent.left {
			uncle := grandparent.right
			if uncle != nil && uncle.red {
				uncle.red = false
				parent.red = false
				grandparent.red = true
				current = grandparent
			} else {
				if current == parent.right {
					current = parent
					tree.RotateLeft(current)
					parent = current.parent
				}
				parent.red = false
				grandparent.red = true
				tree.RotateRight(grandparent)
			}
		}
	}
	tree.root.red = false
}

func (tree *RBTree) RotateLeft(pivotParent *RBNode) {
	if pivotParent == nil || pivotParent.right == nil {
		return
	}
	pivot := pivotParent.right
	pivotParent.right = pivot.left
	if pivot.left != nil {
		pivot.left.parent = pivotParent
	}
	pivot.parent = pivotParent.parent
	if tree.root == pivotParent {
		tree.root = pivot
	} else if pivotParent.parent.left == pivotParent {
		pivotParent.parent.left = pivot
	} else if pivotParent.parent.right == pivotParent {
		pivotParent.parent.right = pivot
	}
	pivot.left = pivotParent
	pivotParent.parent = pivot
}

func (tree *RBTree) RotateRight(pivotParent *RBNode) {
	if pivotParent == nil || pivotParent.left == nil {
		return
	}
	pivot := pivotParent.left
	pivotParent.left = pivot.right
	if pivot.right != nil {
		pivot.right.parent = pivotParent
	}
	pivot.parent = pivotParent.parent
	if tree.root == pivotParent {
		tree.root = pivot
	} else if pivotParent.parent.left == pivotParent {
		pivotParent.parent.left = pivot
	} else if pivotParent.parent.right == pivotParent {
		pivotParent.parent.right = pivot
	}
	pivot.right = pivotParent
	pivotParent.parent = pivot
}

func (tree *RBTree) Display() {
	tree.display(tree.root, "", true)
}

func (tree *RBTree) display(node *RBNode, prefix string, isTail bool) {
	if node == nil || node == tree.null || node.value == nil {
		return
	}

	if node.right != nil {
		nextPrefix := prefix
		if isTail {
			nextPrefix += "│   "
		} else {
			nextPrefix += "    "
		}
		tree.display(node.right, nextPrefix, false)
	}

	color := "B"
	if node.red {
		color = "R"
	}
	branch := "└── "
	if !isTail {
		branch = "┌── "
	}
	fmt.Printf("%s%s%d(%s)\n", prefix, branch, *node.value, color)

	if node.left != nil {
		nextPrefix := prefix
		if isTail {
			nextPrefix += "    "
		} else {
			nextPrefix += "│   "
		}
		tree.display(node.left, nextPrefix, true)
	}
}
