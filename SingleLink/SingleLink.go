package SingleLink

import (
	"fmt"
)

//单链表的接口
type SingleLink interface {
	//增删查改
	GetFirstNode() *Node                                     //返回第一个数据节点
	InsertNodeFront(node *Node)                              //头部插入节点
	InsertNodeBack(node *Node)                               //尾部插入节点
	InsertNodeFrontValue(value interface{}, node *Node) bool //在节点的前面插入
	InsertNodeBackValue(value interface{}, node *Node) bool  //在节点的后面插入
	GetNodeAtIndex(index int) *Node                          //根据索引返回指定位置节点
	DeleteNode(node *Node) bool                              //删除一个节点
	DeleteNodeAtIndex(index int) bool                        //根据索引删除节点
	String() string
	Length() int
	GetHead() *Node

	//返回链表字符串
	ReverseList()                          //迭代反转链表，返回头节点
}

//链表的结构
type LinkList struct {
	head   *Node //链表的头指针
	length int   //链表的长度
}

//创建链表:初始化链表的头指针
func NewLinkList() *LinkList {
	head := NewSingleNode(nil)
	return &LinkList{
		head:   head,
		length: 0,
	}
}

func (list *LinkList) GetFirstNode() *Node {
	return list.head.PNext
}

//头部插入
func (list *LinkList) InsertNodeFront(node *Node) {
	//如果头指针指向的为空
	if list.head == nil {
		list.head.PNext = node
		node.PNext = nil
		//这里确实应该加1 因为头指针指向的为空，直接就在后面拼接
		list.length++
	} else {
		p := list.head //头部备份
		node.PNext = p.PNext
		p.PNext = node
		list.length++
	}
}

//尾部插入
func (list *LinkList) InsertNodeBack(node *Node) {
	//如果头指针指向的为空
	if list.head == nil {
		list.head.PNext = node
		node.PNext = nil
		//这里确实应该加1 因为头指针指向的为空，直接就在后面拼接
		list.length++
	} else {
		p := list.head
		for p.PNext != nil {
			p = p.PNext
		}
		p.PNext = node
		list.length++
	}
}

func (list *LinkList) InsertNodeFrontValue(value interface{}, node *Node) bool {
	//是否找到了这个节点
	isFind := false
	//遍历
	p := list.head
	for p.PNext != nil {
		if p.PNext.Value == value {
			isFind = true
			break
		}
		p = p.PNext
	}
	if isFind {
		node.PNext = p.PNext
		p.PNext = node
		list.length++
		return isFind
	}
	return isFind
}

func (list *LinkList) InsertNodeBackValue(value interface{}, node *Node) bool {
	isFind := false
	p := list.head
	for p.PNext != nil {
		if p.PNext.Value == value {
			isFind = true
			break
		}
		p = p.PNext
	}
	if isFind {
		node.PNext = p.PNext.PNext
		p.PNext.PNext = node
		list.length++
		return isFind
	}
	return isFind
}

func (list *LinkList) GetNodeAtIndex(index int) *Node {
	if index >= list.length || index < -1 {
		return nil
	}
	p := list.head
	for index > -1 {
		//向后循环,然后我一直往后赋值，直到index -1 小于 0 就停止
		p = p.PNext
		index--
	}
	return p
}

//删除节点
func (list *LinkList) DeleteNode(node *Node) bool {
	if node == nil {
		return false
	}
	p := list.head
	for p.PNext != nil {
		if p.PNext == node {
			//逻辑处理
			p.PNext = p.PNext.PNext
			list.length--
			break
		}
		p = p.PNext
	}
	return true
}

//根据索引删除节点
func (list *LinkList) DeleteNodeAtIndex(index int) bool {
	if index >= list.length || index < 0 {
		return false
	}
	p := list.head
	//这里就必须聪明一点,在目标索引(节点)前，停止遍历，然后进行next重新拼接
	for index > 0 {
		p = p.PNext
		index--
	}
	p.PNext = p.PNext.PNext
	list.length--
	return true

}

func (list *LinkList) ReverseList() {
	//链表为空,链表只有一个节点 直接返回
	if list.head == nil || list.head.PNext == nil {
		return
	}
	var pre *Node
	cur := list.head.PNext
	//节点不为空
	for cur != nil {
		temp := cur.PNext
		cur.PNext = pre
		pre = cur
		cur = temp
	}
	list.head.PNext = pre
}

func (list *LinkList) String() string {
	var listString string
	p := list.head
	//有些时候代码真的不要去纠结很多,代码可以就当成一串文字
	for p.PNext != nil {
		listString += fmt.Sprintf("%v-->", p.PNext.Value)
		p = p.PNext
	}
	listString += fmt.Sprintf("nil")
	return listString
}

func (list *LinkList) Length() int {
	return list.length
}

func (list *LinkList) GetHead() *Node {
	return list.head
}
