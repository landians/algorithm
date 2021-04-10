package stack

// Stack 栈的实现
type Stack struct {
	Values []interface{}
	size   int
}

func New(size int) *Stack {
	return &Stack{
		Values: make([]interface{}, 0, size),
	}
}

// Push 入栈
func (s *Stack) Push(v interface{}) {
	if s.IsFull() {
		return
	}
	s.Values = append(s.Values, v)
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return v
}

// IsFull 判断栈满
func (s *Stack) IsFull() bool {
	return cap(s.Values) == len(s.Values)
}

// IsEmpty 判断栈空
func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

// Top 返回栈顶元素
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Values[len(s.Values)-1]
}
