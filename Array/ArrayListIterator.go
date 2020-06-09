package Array

import "errors"

//实现数组迭代器
//为了访问数组方便

type Iterator interface {
	HasNext() bool              //是否有下一个
	Next() (interface{}, error) //下一个
	Remove()                    //删除
	GetIndex() int              //获取索引
}

//构造指针访问数组
type ListIterator struct {
	list         *ArrayList //数组指针
	currentIndex int        //当前索引
}

//初始化迭代器
func (list *ArrayList) NewListIterator() Iterator {
	it := new(ListIterator)
	it.list = list
	it.currentIndex = 0
	return it
}

//是否有下一个
func (it *ListIterator) HasNext() bool {
	return it.currentIndex < it.list.theSize
}

//下一个
func (it *ListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("not has next")
	}
	data, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return data, err
}

//删除
func (it *ListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

//获取当前索引
func (it *ListIterator) GetIndex() int {
	return it.currentIndex
}
