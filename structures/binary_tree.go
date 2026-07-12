package structures

type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	val   *int
}

func (node *BSTNode) Insert(val int) {
	if node.val == nil {
		node.val = &val
		return
	}
	if node.val == &val {
		return
	}
	if val > *node.val {
		if node.right != nil {
			node.right.Insert(val)
			return
		}
		node.right = &BSTNode{val: &val}
		return
	}
	if node.left != nil {
		node.left.Insert(val)
		return
	}
	node.left = &BSTNode{val: &val}
	return
}

func (node *BSTNode) GetMin() int {
	result := node
	for result.left != nil {
		result = result.left
	}
	return *result.val
}

func (node *BSTNode) GetMax() int {
	result := node
	for result.right != nil {
		result = result.right
	}
	return *result.val
}

func (node *BSTNode) Delete(target int) *BSTNode {
	if node.val == nil {
		return nil
	}
	if target < *node.val {
		if node.left != nil {
			node.left = node.left.Delete(target)
		}
		return node
	}
	if target > *node.val {
		if node.right != nil {
			node.right = node.right.Delete(target)
		}
		return node
	}
	if target == *node.val {
		if node.right == nil {
			return node.left
		} else if node.left == nil {
			return node.right
		}
		if node.right != nil && node.left != nil {
			minimum := node.left.GetMin()
			node.val = &minimum
			node.left = node.left.Delete(minimum)
			return node
		}
	}
	return nil
}

func (node *BSTNode) Preorder(visited []int) []int {
	visited = append(visited, *node.val)
	if node.left != nil {
		visited = node.left.Preorder(visited)
	}
	if node.right != nil {
		visited = node.right.Preorder(visited)
	}
	return visited
}

func (node *BSTNode) Postorder(visited []int) []int {
	if node.left != nil {
		visited = node.left.Postorder(visited)
	}
	if node.right != nil {
		visited = node.right.Postorder(visited)
	}
	visited = append(visited, *node.val)
	return visited
}

func (node *BSTNode) Inorder(visited []int) []int {
	if node.left != nil {
		visited = node.left.Inorder(visited)
	}
	visited = append(visited, *node.val)
	if node.right != nil {
		visited = node.right.Inorder(visited)
	}
	return visited
}

func (node *BSTNode) Exists(target int) bool {
	if node.val == nil {
		return false
	}
	if *node.val == target {
		return true
	}
	if *node.val < target && node.right != nil {
		return node.right.Exists(target)
	}
	if node.left != nil {
		return node.left.Exists(target)
	}
	return false
}

func (node *BSTNode) Height() int {
	if node.val == nil {
		return 0
	}
	leftHeight := 0
	rightHeight := 0
	if node.left != nil {
		leftHeight = node.left.Height()
	}
	if node.right != nil {
		rightHeight = node.right.Height()
	}
	return max(leftHeight, rightHeight) + 1
}
