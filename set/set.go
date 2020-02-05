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

// 使用slice实现一个集合
type SliceSet struct {
	data []int
	size int
}

func (s *SliceSet) Init() {

}
func (s *SliceSet) GetSize() int {
	return len(s.data)
}
func (s *SliceSet) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *SliceSet) Add(e int) {
	// 先判断元素是否存在
	for _, d := range s.data {
		if e == d {
			return
		}
	}

	s.data = append(s.data, e)
}

func (s *SliceSet) Remove(e int) {
	// 确定删除位置
	var index int
	for i, d := range s.data {
		if d == e {
			index = i
		}
	}

	s.data = append(s.data[:index], s.data[index+1:]...)
}

func (s *SliceSet) Contains(e int) bool {
	for _, d := range s.data {
		if d == e {
			return true
		}
	}

	return false
}
