package main

import (
	. "DataStructuresAndAlgorithms/SingleLink"
	"fmt"
)

func main() {
	//初始化链表
	list := NewLinkList()
	//构造一个节点
	node1 := NewSingleNode(1)
	//将这个节点插入
	//list.InsertNodeFront(node1)
	list.InsertNodeBack(node1)
	fmt.Println(list)

	node2 := NewSingleNode(2)
	//list.InsertNodeFront(node2)
	list.InsertNodeBack(node2)
	fmt.Println(list)

	node3 := NewSingleNode(3)
	//list.InsertNodeFront(node3)
	list.InsertNodeBack(node3)
	fmt.Println(list)

	node4 := NewSingleNode(4)
	//前插
	//list.InsertNodeFrontValue(2, node4)

	//后插
	list.InsertNodeBackValue(3, node4)
	fmt.Println(list)

	//根据索引取节点
	fmt.Println(list.GetNodeAtIndex(3))

	//删除节点
	//list.DeleteNode(node4)

	//通过索引删除节点
	//list.DeleteNodeAtIndex(2)
	fmt.Println(list)

	//反转链表
	//list.ReverseList()
	//递归反转单链表
	//P := RecursiveReverseList(list.GetHead().PNext)

	//递归反转链表的前N个节点
	p := RecursiveReverseNNode(list.GetHead().PNext, 2)
	fmt.Println(getString(p))
}

func getString(node *Node) string {
	var listString string
	for node != nil {
		listString += fmt.Sprintf("%v-->", node.Value)
		node = node.PNext
	}
	listString += fmt.Sprintf("nil")
	return listString
}
