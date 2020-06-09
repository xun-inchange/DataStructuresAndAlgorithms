package StackArray

//实现接口(定义栈的方法)
type StackArray interface {
	Clear()           //清空栈
	Pop() interface{} //出栈
	Size() int        //返回栈的大小
	Push(interface{}) //入栈
	IsFull() bool     //是否满了
	IsEmpty() bool    //是否为空
}

//结构
type Stack struct {
	dataSource  []interface{} //存储数据
	capSize     int           //最大范围
	currentSize int           //实际使用大小
}

//生成一个栈
func NewStack() *Stack {
	return &Stack{
		dataSource:  make([]interface{}, 0, 10),
		capSize:     10,
		currentSize: 0,
	}
}

//go是自动回收的，所以重新开辟内存那么就清除了
func (s *Stack) Clear() {
	s.dataSource = make([]interface{}, 0, 10)
	s.capSize = 10
	s.currentSize = 0
}

//弹出最后一个数字
func (s *Stack) Pop() interface{} {
	//为空不能弹出
	if s.IsEmpty() {
		return nil
	}
	data := s.dataSource[s.currentSize-1]
	s.currentSize--
	return data
}

func (s *Stack) Size() int {
	return s.currentSize
}

func (s *Stack) Push(data interface{}) {
	//入栈的时候，满了也不能再入栈
	if s.IsFull() {
		return
	}
	s.dataSource = append(s.dataSource, data)
	s.currentSize++
}

func (s *Stack) IsFull() bool {
	return s.currentSize == s.capSize
}

func (s *Stack) IsEmpty() bool {
	return s.currentSize == 0
}
