package BinaryTree

type Node struct {
	data  int
	left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node //二叉树的根节点
	Size int   //二叉树的数据的数量
}
