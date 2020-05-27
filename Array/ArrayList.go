package Array

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	dataStore []interface{} //数组储存
	theSize   int           //数组大小
}

type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //抓取第几个元素
	Set(index int, newVal interface{}) error    //修改数据
	Insert(index int, newVal interface{}) error //插入数据
	Append(newVal interface{})                  //添加数据
	Delete(index int) error                     //删除数据
	Clear()                                     //清空数组
	String() string                             //返回字符串
	checkIsFull()                               //开辟内存
}

//实例化ArrayList
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	//初始化一个长度为0，容量为10的切片
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) checkIsFull() {
	if list.theSize == cap(list.dataStore) {
		//make 中间的参数为0表示没有开辟内存，容量是允许你使用这么多内存
		newStore := make([]interface{}, 2*list.theSize, 2*list.theSize)
		copy(newStore, list.dataStore)
		list.dataStore = newStore
	}
}

//获取数组大小
func (list *ArrayList) Size() int {
	return list.theSize
}

//根据索引获取值
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index >= list.theSize || index < 0 {
		return nil, errors.New("索引越界！")
	}
	return list.dataStore[index], nil
}

//根据索引修改数组的值
func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index >= list.theSize || index < 0 {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newVal
	return nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.theSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

//向数组插入数据
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index >= list.theSize || index < 0 {
		return errors.New("索引越界")
	}
	list.checkIsFull()                               //监测内存，如果满了就自动追加
	list.dataStore = list.dataStore[:list.theSize+1] //插入的话，自动把内存往后扩充一位
	//插入数据,数组中的数据从后往前移动
	for i := list.theSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal //插入数据
	list.theSize++
	return nil
}

//根据索引删除数据
func (list *ArrayList) Delete(index int) error {
	if index >= list.theSize || index < 0 {
		return errors.New("索引越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.theSize--
	return nil
}

//Golang有自动回收垃圾的机制，索引可以重新指定引用的底层数组，这样子之前的数据因为没有指定它的指针，所以会被自动回收，清空内存
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}
