package main

import (
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

func main() {
	testBst()
	//testStack()
}

func testBst() {
	var t *tree.BinarySearchTree
	root := t.Init()

	fmt.Println(root.GetSize())

	root.Insert(10)
	root.Insert(20)
	root.Insert(30)
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
}
