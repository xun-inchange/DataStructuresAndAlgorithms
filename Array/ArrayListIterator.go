package Array

//实现数组迭代器
//为了访问数组方便

type Iterator interface {
	HasNext() bool              //是否有下一个
	Next() (interface{}, error) //下一个
	Remove()                    //删除
	GetIndex() int              //获取索引
}

type Iterable interface {
}

//构造指针访问数组
type ListIterator struct {
	list         *ArrayList //数组指针
	currentIndex int        //当前索引
}
