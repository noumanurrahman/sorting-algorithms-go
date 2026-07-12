package examples

import (
	"fmt"
	"sorting-algorithms/structures"
)

func BinaryTree() {
	bstNode := new(structures.BSTNode)
	bstNode.Insert(5)
	bstNode.Insert(3)
	bstNode.Insert(7)
	bstNode.Insert(4)
	bstNode.Insert(6)
	bstNode.Insert(8)
	fmt.Println(bstNode.GetMin())
	fmt.Println(bstNode.GetMax())
	preorder := bstNode.Preorder([]int{})
	postorder := bstNode.Postorder([]int{})
	inorder := bstNode.Inorder([]int{})
	fmt.Println(preorder, postorder)
	fmt.Println(inorder)
	fmt.Println(bstNode.Exists(7))
	bstNode.Delete(7)
	fmt.Println(bstNode.Exists(7))
	fmt.Println(bstNode.Height())
}
