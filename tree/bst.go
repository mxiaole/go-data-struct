package tree

import (
	"data-struct/queue"
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

// 二叉搜索树的层序遍历
func (bst *BinarySearchTree) LevelOrder(root *Node) {
	// 借助队列实现层序遍历
	var q queue.Queue
	q.Init()

	// 根节点入队列
	q.Push(root)

	for !q.IsEmpty() {
		// 根节点出队列，然后左右子树的根节点分别入队列
		res := q.Pop().(*Node)
		fmt.Println(res.Data)
		if res.Left != nil {
			q.Push(res.Left)
		}
		if res.Right != nil {
			q.Push(res.Right)
		}
	}

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

// 查找树中的最小值
func (bst *BinarySearchTree) Minimum(root *Node) *Node {
	// 找到最小值
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}

	return cur
}

// 查找树中的最大值
func (bst *BinarySearchTree) Maximum(root *Node) *Node {
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	return cur
}

// 删除二分搜索树中的最小值
func (bst *BinarySearchTree) RemoveMin(root *Node) *Node {
	// 如果删除的节点的根节点的左子树为空，那么这个根节点就是要删除的节点
	if root.Left == nil {
		right := root.Right // 先保存它的右子树
		bst.Size--
		return right
	}
	root.Left = bst.RemoveMin(root.Left)
	return root
}

// 删除二分搜索树中的最大值
func (bst *BinarySearchTree) RemoveMax(root *Node) *Node {
	if root.Right == nil {
		left := root.Left
		bst.Size--
		return left
	}

	root.Right = bst.RemoveMax(root.Right)
	return root
}

// 删除二分搜索树的元素e
// root为一个树的跟根节点，返回是删除了指定元素的树的根节点
func (bst *BinarySearchTree) Remove(root *Node, e int) *Node {
	if root == nil {
		return nil
	}
	if root.Data > e { // e在root为根的树的左子树上
		root.Left = bst.Remove(root.Left, e)
		return root
	} else if root.Data < e {
		root.Right = bst.Remove(root.Right, e)
		return root
	}
	// 如果该节点只有左子树
	if root.Right == nil {
		// 保存该节点的左子树
		left := root.Left
		bst.Size--
		return left
	}
	// 如果该节点只有右子树
	if root.Right == nil {
		// 保存该节点的右子树
		right := root.Right
		bst.Size--
		return right
	}

	// 如果该节点既有左子树又有右子树
	// 找到替换该节点的节点, 该节点是它的右子树中的最小值
	node := bst.Minimum(root.Right)
	// 	设置node的右子树
	node.Right = bst.RemoveMin(root.Right) // 替换的节点的右子树为，将root右子树中的最小值的树的根
	// 设置node的左子树
	node.Left = root.Left
	return node
}
