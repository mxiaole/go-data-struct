package tree

import "fmt"

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

// 二叉树的先序遍历
func (bst *BinarySearchTree) PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	bst.PreOrder(root.Left)
	bst.PreOrder(root.Right)
}
