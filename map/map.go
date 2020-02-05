package mymap

type Map interface {
	Init()
	Add(k string, v string)
	Remove(k string) // 给定指定的key，删除value
	Contains(k string) bool
	Get(k string) string
	GetSize() int
	IsEmpty() bool
}

// 数据节点
type node struct {
	k string
	v string
}

// 使用slice实现map
type SliceMap struct {
	data []node // 存放map中元素的slice
	size int    // map中元素的个数
}

func (m *SliceMap) Init() {
	m.data = nil
	m.size = 0
}

func (m *SliceMap) Add(k string, v string) {
	// 判断k是否存在，不存在写入，存在更新
	for _, d := range m.data {
		if k == d.k {
			d.v = v
			return
		}
	}

	var node node
	node.k = k
	node.v = v
	m.data = append(m.data, node)
	m.size++
}

func (m *SliceMap) Remove(k string) {
	var index int
	for i, d := range m.data {
		if d.k == k {
			index = i
			break
		}
	}

	m.data = append(m.data[:index], m.data[index+1:]...)
	m.size--
}

func (m *SliceMap) Contains(k string) bool {
	for _, d := range m.data {
		if d.k == k {
			return true
		}
	}

	return false
}

func (m *SliceMap) Get(k string) string {
	for _, d := range m.data {
		if d.k == k {
			return d.v
		}
	}

	return ""
}

func (m *SliceMap) GetSize() int {
	return m.size
}

func (m *SliceMap) IsEmpty() bool {
	return m.size == 0
}
