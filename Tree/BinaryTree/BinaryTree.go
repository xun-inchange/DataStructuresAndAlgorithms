package BinaryTree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	Root *Node //二叉树的根节点
	Size int   //二叉树的数据的数量
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
		Size: 0,
	}
}

func (bt *BinaryTree) GetSize() int {
	return bt.Size
}

//判断二叉树是否为空
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Size == 0
}

//根节点插入
//插入节点：从根节点开始递归下去然后查找到合适位置进行插入
func (bt *BinaryTree) Add(data int) {
	bt.Root = bt.add(bt.Root, data)
}

//插入节点
func (bt *BinaryTree) add(node *Node, data int) *Node {
	//如果节点是空的，那么就将数据插入当前的节点
	if node == nil {
		bt.Size++
		return &Node{
			data:  data,
			left:  nil,
			right: nil,
		}
	}
	//要插入的节点data比当前的节点data小就往右边插
	//两种情况：1.这个数据比当前数据小往左边插入 2.这个数据比当前的数据大往右边插入
	//如果插入的数据和当前数据一致，也就是说重复了，那么就直接return,不插入了
	if data < node.data {
		node.left = bt.add(node.left, data)
	} else if data > node.data {
		node.right = bt.add(node.right, data)
	}
	//返回根节点
	return node
}

//查找数据是否存在 方法使用:递归
func (bt *BinaryTree) IsIn(data int) bool {
	return bt.isIn(bt.Root, data)
}

func (bt *BinaryTree) isIn(node *Node, data int) bool {
	//如果根节点是空的，那么就直接返回false
	if node == nil {
		return false
	}
	if node.data == data {
		return true
	} else if node.data < data {
		//如果当前节点数据比data小，那么就继续往左边递归
		return bt.isIn(node.left, data)
	} else {
		//如果当前节点数据比data大，那么就继续往右边递归
		return bt.isIn(node.right, data)
	}
}

func (bt *BinaryTree) GetMaximum() int {
	if bt.Root == nil {
		panic("树为空")
	}
	return bt.getMaximum(bt.Root).data
}

//迭代取得二叉树的极大值(极大的节点)
func (bt *BinaryTree) getMaximum(node *Node) *Node {
	if node.right == nil {
		//只要右边的节点不等于nil，就一直递归右边下,否则就返回当前节点作为最终结果
		return node
	}
	return bt.getMaximum(node.right)
}

func (bt *BinaryTree) GetMinimum() int {
	if bt.Root == nil {
		panic("树为空")
	}
	return bt.getMinimum(bt.Root).data
}

func (bt *BinaryTree) getMinimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return bt.getMinimum(node.left)
}

//树的遍历///////////////////

//前序遍历:根->左子树->右子树
func (bt *BinaryTree) PreOrder() {
	bt.preOrder(bt.Root)
}
func (bt *BinaryTree) preOrder(node *Node) {
	//如果节点为空就返回
	if node == nil {
		return
	}
	fmt.Println(node.data)
	bt.preOrder(node.left)
	bt.preOrder(node.right)
}

//中序遍历:左子树->根->右子树
func (bt *BinaryTree) InOrder() {
	bt.inOrder(bt.Root)
}
func (bt *BinaryTree) inOrder(node *Node) {
	if node == nil {
		return
	}
	bt.inOrder(node.left)
	fmt.Println(node.data)
	bt.inOrder(node.right)
}

//后序遍历:左子树->右子树->根
func (bt *BinaryTree) PostOrder() {
	bt.postOrder(bt.Root)
}
func (bt *BinaryTree) postOrder(node *Node) {
	if node == nil {
		return
	}
	bt.postOrder(node.left)
	bt.postOrder(node.right)
	fmt.Println(node.data)
}

//删除二叉树中的最大值和最小值
func (bt *BinaryTree) RemoveMax() {
	bt.Root = bt.removeMax(bt.Root)
}
func (bt *BinaryTree) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		bt.Size--
		return leftNode
	}
	node.right = bt.removeMax(node.right)
	return node
}

func (bt *BinaryTree) RemoveMin() {
	bt.Root = bt.removeMin(bt.Root)
}

func (bt *BinaryTree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		bt.Size--
		return rightNode
	}
	node.left = bt.removeMin(node.left)
	return node
}

//任意删除,指定删除节点中的某个数据
func (bt *BinaryTree) RemoveNode(data int) {
	bt.Root = bt.removeNode(bt.Root, data)
}

func (bt *BinaryTree) removeNode(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.data {
		node.left = bt.removeNode(node.left, data)
		return node
	} else if data > node.data {
		node.right = bt.removeNode(node.right, data)
		return node
	} else {
		//分情况进行讨论
		//1.左子树为nil时
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			bt.Size--
			return rightNode
		} else if node.right == nil {
			//2.右子树为空的时候
			leftNode := node.left
			node.left = nil
			bt.Size--
			return leftNode
		}
		//最后一种情况是左，右子树都不为空的时候
		//先找一个比当前节点大的节点，从右子树找一个最小的节点来代替现在的节点
		tempNode := bt.getMinimum(node.right)
		//然后将目标节点的左右子树接起来
		//这里顺序一定要不能弄错了
		tempNode.right = bt.removeMin(node.right)
		tempNode.left = node.left
		//然后将目标节点指向的左右子树全部删除掉
		node.left = nil
		node.right = nil
		bt.Size--
		return tempNode
	}
}
