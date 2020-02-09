package tree

import "fmt"

// 并查集

type UF interface {
	Init(size int)
	IsConnected(p, q int) bool
	UnionElements(p, q int)
	GetSize() int
}

type UnionFind struct {
	id []int // 数组索引表示元素的编号，数值中的值表示元素所属的集合编号
}

func (uf *UnionFind) Init(size int) {
	uf.id = make([]int, size)
	for i := 0; i < len(uf.id); i++ {
		uf.id[i] = i
	}
}

// 判断两个元素是否属于同一个集合
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// 合并元素p和元素q所在的集合
func (uf *UnionFind) UnionElements(p, q int) {
	// 获取元素p所属的集合的编号
	s1 := uf.find(p)
	// 获取元素q所属的集合的编号
	// 将元素q所属的集合及所有属于该集合的元素都改成该集合
	s2 := uf.find(q)
	for i, v := range uf.id {
		if v == s2 {
			uf.id[i] = s1
		}
	}
}

// 获取并查集中元素的个数
func (uf *UnionFind) GetSize() int {
	return len(uf.id)
}

// 私有方法，查找元素p所属的集合编号
func (uf *UnionFind) find(p int) int {
	if p < 0 || p > len(uf.id) {
		panic("invalid p")
	}
	return uf.id[p]
}

func (uf *UnionFind) Print() {
	fmt.Println(uf.id)
}

type QuickUnion struct {
	id []int
}

func (qu *QuickUnion) Init(size int) {
	qu.id = make([]int, size)
	for i := 0; i < len(qu.id); i++ {
		qu.id[i] = i
	}
}

func (qu *QuickUnion) IsConnected(p, q int) bool {
	// 查找p的根
	root1 := qu.find(p)
	// 查找q的根
	root2 := qu.find(q)

	if root1 == root2 {
		return true
	}

	return false
}

func (qu *QuickUnion) UnionElements(p, q int) {
	root1 := qu.find(p)
	root2 := qu.find(q)

	if root1 == root2 {
		return
	}

	qu.id[root1] = root2
}

func (qu *QuickUnion) GetSize() int {
	return len(qu.id)
}

func (qu *QuickUnion) find(p int) int {
	for p != qu.id[p] {
		p = qu.id[p]
	}

	return p
}

func (qu *QuickUnion) Print() {
	fmt.Println(qu.id)
}
