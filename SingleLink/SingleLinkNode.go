package SingleLink


//链表节点
type Node struct {
	Value interface{}
	PNext *Node
}

//构造一个节点
func NewSingleNode(data interface{}) *Node {
	//返回当前节点的地址
	return &Node{data, nil}
}

//返回数据
func (node *Node) GetValue() interface{} {
	return node.Value
}

//返回一个节点
func (node *Node) GetPNext() *Node {
	return node.PNext
}
