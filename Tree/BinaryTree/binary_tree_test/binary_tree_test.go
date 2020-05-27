package binary_tree_test

import (
	"DataStructuresAndAlgorithms/Tree/BinaryTree"
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := BinaryTree.NewBinaryTree()
	nums := []int{4, 2, 1, 3, 6, 5, 7}
	for _, num := range nums {
		bt.Add(num)
	}
	bt.RemoveNode(4)
	//前序遍历
	fmt.Println("---------------前序遍历----------------------")
	bt.PreOrder()
	//fmt.Println("---------------中序遍历----------------------")
	////中序遍历
	//bt.InOrder()
	//fmt.Println("---------------后序遍历----------------------")
	////后序遍历
	//bt.PostOrder()
}
