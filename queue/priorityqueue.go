package queue

import (
	heap2 "data-struct/heap"
	"fmt"
)

// 使用堆实现一个优先级队列

type PriorityQueue struct {
	data *heap2.SliceHeap // 存放队列中的元素的二叉堆
}

func (q *PriorityQueue) Init() {
	q.data = new(heap2.SliceHeap)
	q.data.Init()
}
func (q *PriorityQueue) Push(e interface{}) {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	if d, ok := e.(int); ok {
		q.data.Add(d)
		return
	}
	panic("type is not supported")
}

func (q *PriorityQueue) Pop() interface{} {
	return q.data.Remove()
}

func (q *PriorityQueue) GetSize() int {
	return q.data.GetSize()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.data.IsEmpty()
}
