package binary_tree_test

import (
	"DataStructuresAndAlgorithms/Tree/BinaryTree"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := BinaryTree.NewBinaryTree()
	nums := []int{4, 2, 1, 3, 6, 5, 7}
	for _, num := range nums {
		bt.Add(num)
	}
	//前序遍历
	bt.PreOrder()
	//中序遍历
	bt.InOrder()
	//后序遍历
	bt.PostOrder()

}
