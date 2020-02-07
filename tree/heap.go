package tree

type SliceHeap struct {
	data []int // 存放堆中元素的slice
}

type Heap interface {
	Init()
	GetSize() int      // 获取堆中元素
	IsEmpty() bool     // 判断堆是否为空
	Add(e int)         // 向堆中添加元素
	Remove() int       // 从堆中取出最大的元素
	Replace(e int)     // 将堆中的最大值，替换成元素e
	Heapify(arr []int) // 将给定的数组初始化为一个堆
}

// 获取index下标的父节点的下标
func parent(index int) int {
	if index == 0 {
		panic("index is invalid")
	}
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func swap(i int, j int, slice []int) {
	temp := slice[i]
	slice[i] = slice[j]
	slice[j] = temp
}

// 将一个指定位置的元素在堆中进行上浮siftUp
func siftUp(index int, heap *SliceHeap) {
	// 如果这个元素的值大于它的父节点的值或者这个元素到达了下标为0的位置即根节点, 这个index下标所对应的节点就需要和父节点交换
	for index > 0 && heap.data[index] > heap.data[parent(index)] {
		swap(index, parent(index), heap.data)
		index = parent(index)
	}
}

// 将指定index的元素进行下沉
func siftDown(index int, heap *SliceHeap) {
	// index的左孩子已经越界
	for leftChild(index) < len(heap.data) {
		// 确定index作用孩子中较大的是哪个
		leftIndex := leftChild(index)     // 获取左孩子的值
		var maxIndex int                  // 左右孩子中较大的值的下标
		if leftIndex+1 < len(heap.data) { // 首先保证右孩子存在
			// 比较右孩子的值和左孩子的值哪个更大
			if heap.data[leftIndex] > heap.data[leftIndex+1] { // 如果左孩子更大
				maxIndex = leftIndex
			} else {
				maxIndex = leftIndex + 1
			}
		} else {
			maxIndex = leftIndex
		}
		// index和maxIndex对应的元素比较
		if heap.data[index] > heap.data[maxIndex] { // 如果index元素比左右孩子的值都大，不用下沉
			break
		}
		// 交换index和maxIndex
		swap(index, maxIndex, heap.data)
		index = maxIndex
	}
}
func (heap *SliceHeap) Init() {
	heap.data = nil
}
func (heap *SliceHeap) GetSize() int {

	return len(heap.data)
}
func (heap *SliceHeap) IsEmpty() bool {
	return len(heap.data) == 0
}

func (heap *SliceHeap) Add(e int) {
	// 将待添加的元素直接添加到数组的末尾
	heap.data = append(heap.data, e)
	// 将这个新添加进来的元素进行sift-up, 维护最大堆的性质
	siftUp(len(heap.data)-1, heap)

}
func (heap *SliceHeap) Remove() int {
	// 堆顶元素的最大值
	ret := heap.data[0]
	// 将堆顶元素和最后一个元素进行交换
	swap(0, len(heap.data)-1, heap.data)
	// 删除最后一个元素
	heap.data = heap.data[:len(heap.data)-1]
	// 将第一个元素进行下沉
	siftDown(0, heap)

	return ret
}

func (heap *SliceHeap) Replace(e int) {
	// 1. 将元素e直接和堆顶元素进行替换
	heap.data[0] = e
	// 2. 将堆顶元素进行siftDown
	siftDown(0, heap)
}

func (heap *SliceHeap) Heapify(arr []int) {
	// 从最后一个非叶子节点（最后一个元素的父节点）开始，从后向前进行sift-down
	for i := parent(len(arr) - 1); i > 0; i-- {
		siftDown(i, heap)
	}
}
