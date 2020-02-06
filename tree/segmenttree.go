package tree

// 线段树, 使用数组实现

type SegmentTree struct {
	data []interface{}
	tree []interface{}
}

// 线段树的基本操作
// 1. 线段树的初始化
func (tree *SegmentTree) Init(array []interface{}) {
	tree.tree = make([]interface{}, 4*len(array))
	tree.data = array
	// 创建线段树
	buildSegmentTree(0, 0, len(array)-1, tree)
}

// 创建线段树的辅助函数
// 给定一个树的根节点的index和要创建线段树的区间[l...r]
func buildSegmentTree(index int, l int, r int, tree *SegmentTree) {
	if l == r { // 如果区间中只有一个元素
		tree.tree[index] = tree.data[l]
		return
	}
	// 左孩子怎么创建
	left := leftChild(index)
	// 右孩子怎么创建
	right := rightChild(index)

	mid := l + (r-l)/2

	buildSegmentTree(left, l, mid, tree)
	buildSegmentTree(right, mid+1, r, tree)

	// 根节点的值是什么？对[l...r]区间中的数据进行处理之后的结果, 此处以整数的求和为例
	tree.tree[index] = tree.tree[left].(int) + tree.tree[right].(int)
}

// 线段树实现的辅助函数, 找一个节点的左右两个孩子节点
func leftChild(index int) int {
	return 2*index + 1
}
func rightChild(index int) int {
	return 2*index + 2
}

// 2. 查询某个区间
func (tree *SegmentTree) Query(l int, r int) interface{} {
	// 判断参数的合法性
	if l > r || l < 0 || r > len(tree.tree) {
		panic("invalid index")
	}
	res := query(0, 0, len(tree.data)-1, l, r, tree)

	return res
}

// 查询的辅助函数
// 输入：查询的线段树的根和这个根的值是哪个线段区间[l ... r], 查询的区间[ql ... qr]
func query(index, l, r, ql, qr int, tree *SegmentTree) interface{} {
	if l == ql && r == qr {
		return tree.tree[index]
	}

	mid := l + (r-l)/2            // 左孩子对应的区间[l, mid]
	leftIndex := leftChild(index) // 右孩子对应的区间[mid+1, r]
	rightIndex := rightChild(index)

	if qr <= mid { // 左孩子对应的区间中
		return query(leftIndex, l, mid, ql, qr, tree)
	} else if ql >= mid+1 {
		return query(rightIndex, mid+1, r, ql, qr, tree)
	}
	// 左右子树中都有
	leftResult := query(leftIndex, l, mid, ql, mid, tree)
	rightResult := query(rightIndex, mid+1, r, mid+1, qr, tree)

	return leftResult.(int) + rightResult.(int)
}
