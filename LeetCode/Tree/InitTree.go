package Tree

import "fmt"

type BinaryTree struct {
	Root *Node
	Size int
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
		Size: 0,
	}
}

func (bt *BinaryTree) PreOrder() {
	bt.preOrder(bt.Root)
}
func (bt *BinaryTree) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	bt.preOrder(node.Left)
	bt.preOrder(node.Right)
}
