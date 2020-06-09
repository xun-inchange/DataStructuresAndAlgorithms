package Array

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
	MyArray *ArrayList //存储数据
	capSize int        //最大范围
}

//生成一个栈
func NewArrayListStack() *Stack {
	return &Stack{
		MyArray: NewArrayList(),
		capSize: 10,
	}
}

//go是自动回收的，所以重新开辟内存那么就清除了
func (s *Stack) Clear() {
	s.MyArray.Clear()
	s.capSize = 10
}

//弹出最后一个数字
func (s *Stack) Pop() interface{} {
	//为空不能弹出
	if s.IsEmpty() {
		return nil
	}
	data := s.MyArray.dataStore[s.MyArray.theSize-1]
	s.MyArray.Delete(s.MyArray.theSize - 1)
	return data
}

func (s *Stack) Size() int {
	return s.MyArray.theSize
}

func (s *Stack) Push(data interface{}) {
	//入栈的时候，满了也不能再入栈
	if s.IsFull() {
		return
	}
	s.MyArray.Append(data)
}

func (s *Stack) IsFull() bool {
	return s.MyArray.theSize == s.capSize
}

func (s *Stack) IsEmpty() bool {
	return s.MyArray.theSize == 0
}
