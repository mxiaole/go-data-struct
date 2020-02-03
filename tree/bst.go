package tree

import (
	"data-struct/stack"
	"fmt"
)

/**
二叉搜索树
1. 节点的定义
*/

// 1. 节点的定义
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// 2. 二叉树的定义
type BinarySearchTree struct {
	Root *Node // 树的根节点
	Size int   // 树中节点的个数
}

// 3. 实现二叉树的常用操作

// 初始化二叉搜索树
func (bst *BinarySearchTree) Init() *BinarySearchTree {
	return &BinarySearchTree{nil, 0}
}

// 获取树中节点的个数
func (bst *BinarySearchTree) GetSize() int {
	return bst.Size
}

// 判断当前的二分搜索树是否为空
func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.Size == 0
}

// 向二分搜索树中添加节点
func (bst *BinarySearchTree) Insert(e int) {
	bst.Root = add(bst.Root, e)
	bst.Size++
}

// Insert方法的辅助方法：
// 输入：一个二叉树的根节点和待插入的元素e
// 输出: 二分搜索树的根节点
func add(root *Node, e int) *Node {
	// 如果根节点为空，那么直接创建给节点就可以
	if root == nil {
		return &Node{Data: e, Left: nil, Right: nil}
	}
	// 如果根节点不是空，就需要判断元素e是放在左孩子还是右孩子
	if e < root.Data { // 如果元素e大于左孩子
		root.Left = add(root.Left, e)
	} else if e > root.Data {
		root.Right = add(root.Right, e)
	}

	// 如果元素e和当前的root.data值相同就忽略
	return root
}

// 判断元素e是否在树中
func (bst *BinarySearchTree) Contains(e int) bool {
	return contains(bst.Root, e)
}

// Contains的辅助方法
func contains(root *Node, e int) bool {
	if root == nil {
		return false
	}

	if root.Data == e {
		return true
	} else if root.Data > e {
		return contains(root.Left, e)
	} else {
		return contains(root.Right, e)
	}
}

// 二叉搜索树的先序遍历
func (bst *BinarySearchTree) PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	bst.PreOrder(root.Left)
	bst.PreOrder(root.Right)
}

// 二叉搜索树的中序遍历
func (bst *BinarySearchTree) MidOrder(root *Node) {
	if root == nil {
		return
	}
	bst.MidOrder(root.Left)
	fmt.Println(root.Data)
	bst.MidOrder(root.Right)
}

// 二叉搜索树的后序遍历
func (bst *BinarySearchTree) PostOrder(root *Node) {
	if root == nil {
		return
	}
	bst.PostOrder(root.Left)
	bst.PostOrder(root.Right)
	fmt.Println(root.Data)
}

// 前序遍历的非递归实现
// 借助一个栈来实现：首先将树的根节点放入栈中，然后根节点出栈，同时根节点的左右节点依次入栈
func (bst *BinarySearchTree) PreOrderNoneRecursive(root *Node) {
	// 定义一个栈
	var s stack.Stack
	s.Init()
	// root入栈
	s.Push(root)

	for !s.IsEmpty() {
		res := s.Pop().(*Node)
		fmt.Println(res.Data)
		// root右、左节点依次入栈
		if res.Right != nil {
			s.Push(res.Right)
		}
		if res.Left != nil {
			s.Push(res.Left)
		}
	}
}
