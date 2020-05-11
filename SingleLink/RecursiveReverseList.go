package SingleLink

//递归反转链表的各种实现

//递归反转链表
//就是希望当前节点指向的下一个节点指向的next 指向它本身
//注意的是要修改当前节点了，当前节点指向的下个节点就必须是nil就行了
func RecursiveReverseList(head *Node) *Node {
	//链表为空或者链表的下一个节点为nil就直接返回当前节点
	if head.PNext == nil {
		return head
	}
	P := RecursiveReverseList(head.PNext)
	//当前节点指向的下个节点指向的下个节点指向当前节点,这样就完成了反转
	head.PNext.PNext = head
	head.PNext = nil
	return P
}

var temp *Node

//反转前N个节点
func RecursiveReverseNNode(head *Node, n int) *Node {
	//如果n==1就直接返回，返回本身
	if n == 1 {
		//后置节点
		temp = head.PNext
		return head
	}
	last := RecursiveReverseNNode(head.PNext, n-1)
	head.PNext.PNext = head
	head.PNext = temp
	return last
}

// 反转 m,n之间的节点(反转部分链表)
func RecursiveReversePartList(head *Node, m, n int) *Node {
	//如果m==1就是反转链表从head开始后的n个元素
	if m == 1 {
		return RecursiveReverseNNode(head, n)
	}
	//否则的话就迭代到m==1那里，这个时候n也必须减1,因为m和n之间的间隔不能改变
	head.PNext = RecursiveReversePartList(head.PNext, m-1, n-1)
	return head
}
