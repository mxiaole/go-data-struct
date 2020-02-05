package set

import (
	"data-struct/tree"
)

type Set interface {
	Init()               // 集合初始化
	Add(e int)           // 向集合中添加元素
	Remove(e int)        // 从集合中移除元素
	Contains(e int) bool // 判断集合中是否包含元素
	GetSize() int        // 获取集合中的元素个数
	IsEmpty() bool       // 判断集合是否为空
}

// 使用二叉搜索树实现一个集合
type BstSet struct {
	tree *tree.BinarySearchTree // 保存集合中元素的二叉搜索树
}

// 初始化set
func (s *BstSet) Init() {
	s.tree = s.tree.Init()
}
func (s *BstSet) GetSize() int {
	return s.tree.GetSize()
}

func (s *BstSet) IsEmpty() bool {
	return s.tree.IsEmpty()
}

func (s *BstSet) Add(e int) {
	s.tree.Insert(e)
}
func (s *BstSet) Remove(e int) {
	s.tree.Root = s.tree.Remove(s.tree.Root, e)
}

func (s *BstSet) Contains(e int) bool {
	return s.tree.Contains(e)
}
