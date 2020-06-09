package Array

//实现接口(定义栈的方法)
type StackArrayX interface {
	Clear()           //清空栈
	Pop() interface{} //出栈
	Size() int        //返回栈的大小
	Push(interface{}) //入栈
	IsFull() bool     //是否满了
	IsEmpty() bool    //是否为空
}

//结构
type StackX struct {
	MyArray *ArrayList //存储数据
	MyIt    Iterator
}

//生成一个栈
func NewArrayListStackX() *StackX {
	stackX := new(StackX)
	stackX.MyArray = NewArrayList()
	stackX.MyIt = stackX.MyArray.NewListIterator()
	return stackX
}

//go是自动回收的，所以重新开辟内存那么就清除了
func (s *StackX) ClearX() {
	s.MyArray.Clear()
}

//弹出最后一个数字
func (s *StackX) PopX() interface{} {
	//为空不能弹出
	if s.IsEmptyX() {
		return nil
	}
	data := s.MyArray.dataStore[s.MyArray.theSize-1]
	s.MyArray.Delete(s.MyArray.theSize - 1)
	return data
}

func (s *StackX) SizeX() int {
	return s.MyArray.theSize
}

func (s *StackX) PushX(data interface{}) {
	//入栈的时候，满了也不能再入栈
	if s.IsFullX() {
		return
	}
	s.MyArray.Append(data)
}

func (s *StackX) IsFullX() bool {
	return s.MyArray.theSize >= 10
}

func (s *StackX) IsEmptyX() bool {
	return s.MyArray.theSize == 0
}
