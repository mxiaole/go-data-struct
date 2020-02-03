package stack

// 使用slice实现栈

type Stack struct {
	Data []interface{} // 存放栈中元素的slice
	Size int           // 栈中元素的个数
}

// 栈的基本操作
// 0. 栈的初始化
func (s *Stack) Init() {
	s.Data = []interface{}{}
	s.Size = 0
}

// 1. 元素入栈
func (s *Stack) Push(e interface{}) {
	s.Data = append(s.Data, e)
	s.Size++
}

// 2. 元素出栈
func (s *Stack) Pop() interface{} {
	res := s.Data[s.Size-1]
	s.Size--
	s.Data = s.Data[:s.Size]

	return res
}

// 3. 获取栈中元素个数
func (s *Stack) GetSize() int {
	return s.Size
}

// 4. 判断栈空
func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}
