package main

import (
	"data-struct/graph"
	heap2 "data-struct/heap"
	mymap "data-struct/map"
	"data-struct/queue"
	"data-struct/set"
	"data-struct/stack"
	"data-struct/tree"
	"fmt"
)

func testStack() {
	var s stack.Stack

	// 初始化
	s.Init()
	// 向栈中添加元素
	s.Push(2)
	s.Push(3)
	s.Push(20)
	// 从栈中弹出元素
	r := s.Pop()
	fmt.Println(r)
	r = s.Pop()
	fmt.Println(r)

}

// bst测试
func testBst() {
	var t *tree.BinarySearchTree
	root := t.Init()

	fmt.Println(root.GetSize())

	root.Insert(9)
	root.Insert(7)
	root.Insert(13)
	root.Insert(8)
	root.Insert(10)
	root.Insert(15)
	fmt.Println(root.GetSize())
	fmt.Println("前序遍历结果")
	root.PreOrder(root.Root)
	fmt.Println("中序遍历结果")
	root.MidOrder(root.Root)
	fmt.Println("后序遍历结果")
	root.PostOrder(root.Root)
	fmt.Println(root.IsEmpty())

	fmt.Println(root.Contains(15))
	fmt.Println(root.Contains(34))

	// 非递归实现的前序遍历
	fmt.Println("-------------", "非递归前序遍历")
	root.PreOrderNoneRecursive(root.Root)
	// 层序遍历
	fmt.Println("-------------  层序遍历  -----------------")
	root.LevelOrder(root.Root)
	// 查找最小值
	fmt.Println("最小值", root.Minimum(root.Root))
	// 查找最大值
	fmt.Println("最大值", root.Maximum(root.Root))
	// 删除最小值
	root.Root = root.RemoveMin(root.Root)
	// 删除之后层序遍历
	fmt.Println("删除最小值之后的层序遍历结果")
	root.LevelOrder(root.Root)
	// 删除最大值
	root.Root = root.RemoveMax(root.Root)
	// 删除之后层序遍历
	fmt.Println("删除最大值之后的层序遍历结果")
	root.LevelOrder(root.Root)

	// 删除树中的值为9的节点
	fmt.Println("删除了值为9的节点层序遍历结果")
	root.Root = root.Remove(root.Root, 9)
	root.LevelOrder(root.Root)
}

//set测试
func testSet() {
	//var s set.BstSet
	var s set.SliceSet
	s.Init()
	s.Add(10)
	s.Add(10)
	s.Add(10)
	s.Add(10)
	s.Add(10)
	s.Add(20)
	s.Add(20)
	fmt.Println("集合中元素的个数是: ", s.GetSize())
	fmt.Println("集合中是否包含元素20", s.Contains(20))
	s.Remove(10)
	fmt.Println("集合中元素的个数是: ", s.GetSize())
	fmt.Println("集合中是否包含元素10", s.Contains(10))

}

// map测试
func testMap() {
	var m mymap.SliceMap
	m.Init()

	m.Add("name", "cody")
	fmt.Println(m.Get("name"))
	fmt.Println(m.GetSize())
	m.Add("sex", "man")
	fmt.Println(m.GetSize())
	m.Remove("name")
	fmt.Println(m.Get("cody"))
	fmt.Println(m.GetSize())
}

// heap测试
func testHeap() {
	var heap heap2.SliceHeap
	heap.Init()
	heap.Add(1)
	heap.Add(10)
	heap.Add(2)
	heap.Add(8)
	heap.Add(9)
	heap.Add(5)

	fmt.Println(heap.GetSize())
	fmt.Println(heap.Remove())
}

func testPriorityQueue() {
	var q queue.PriorityQueue
	q.Init()
	q.Push(10)
	q.Push(20)
	q.Push(30)
	q.Push(15)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}

// 线段树测试
func testSegmentTree() {
	var st tree.SegmentTree
	var array = []interface{}{1, 3, 10, 20, 10}
	st.Init(array)

	fmt.Println(st.Query(1, 2))

	st.Update(2, 3)
	fmt.Println(st.Query(2, 3))
}

// 图测试
func testGraph() {
	var g graph.AdjMatrix
	g.Init("graph/graph.txt")
	var g2 graph.AdjList
	g2.Init("graph/graph.txt")
	g2.Print()
	fmt.Println()
	// 判断0和3是否有边
	fmt.Println(g.HasEdge(0, 2))
	// 获取顶点3的相邻的节点
	fmt.Println(g.Adj(2))
	// 获取顶点3的度
	fmt.Println(g.Degree(2))
}

// 字典树测试
func testTrie() {
	var tree *tree.Trie
	tree.Init()
	fmt.Println(tree.GetSize())
}
func main() {
	//testBst()
	//testStack()
	//testSet()
	//testMap()
	//testHeap()
	//testPriorityQueue()
	//testSegmentTree()
	//testGraph()
	//testTrie()
}
