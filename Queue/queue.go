package Queue

type MyQueue interface {
	Size() int                //获取队列大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //是否为空
	EnQueue(data interface{}) //入队
	DeQueue() interface{}     //出队
	Clear()                   //清空
}
type Queue struct {
	dataStore []interface{}
	size      int
}

//生成一个队列
func NewQueue() *Queue {
	return &Queue{
		dataStore: make([]interface{}, 0),
		size:      0,
	}
}

func (q *Queue) Size() int {
	return q.size
}

//返回第一个数据
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

//返回最后一个元素
func (q *Queue) End() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.dataStore[q.size-1]
}

//入队
func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data)
	q.size++
}

//出队
func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.dataStore[0]
	if q.size > 1 {
		q.dataStore = q.dataStore[1:q.size]
	} else {
		q.dataStore = make([]interface{}, 10) //内存不够了，重新开辟内存
	}
	q.size--
	return data
}
