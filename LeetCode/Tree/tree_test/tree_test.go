package tree_test

import (
	"DataStructuresAndAlgorithms/LeetCode/Tree"
	"testing"
)

func TestTree(t *testing.T) {
	bt := Tree.NewTree()
	bt.Root = &Tree.Node{
		Val: 1,
	}
	bt.Root.Left = &Tree.Node{
		Val: 2,
	}
	bt.Root.Right = &Tree.Node{
		Val: 2,
	}
	bt.Root.Left.Left = &Tree.Node{
		Val: 3,
	}
	bt.Root.Left.Right = &Tree.Node{
		Val: 4,
	}
	bt.Root.Right.Left = &Tree.Node{
		Val: 4,
	}
	bt.Root.Right.Right = &Tree.Node{
		Val: 3,
	}
	bt.PreOrder()
}
