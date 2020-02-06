package queue

type Q interface {
	Init()
	Push(e interface{})
	Pop() interface{}
	GetSize() int
	IsEmpty() bool
}

// 使用slice实现queue

// 定义队列
type Queue struct {
	Data []interface{} // 存储队列中元素的slice
	Size int           // 队列中元素的个数
}

// 定义队列的基本操作
// 0. 队列的初始化
func (q *Queue) Init() {
	q.Data = []interface{}{}
	q.Size = 0
}

// 1. 元素入队列
func (q *Queue) Push(e interface{}) {
	q.Data = append(q.Data, e)
	q.Size++
}

// 2. 元素出队列
func (q *Queue) Pop() interface{} {
	res := q.Data[0]
	q.Size--
	q.Data = q.Data[1:]

	return res
}

// 3. 获取队列中元素的个数
func (q *Queue) GetSize() int {
	return q.Size
}

// 4. 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
