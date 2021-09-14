package stack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New(10)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	fmt.Println("===== 入栈 =====")
	fmt.Println(s.Values)
	fmt.Println("===== 入栈 =====")

	fmt.Println()

	s.Pop()
	s.Pop()
	s.Pop()

	fmt.Println("===== 出栈 =====")
	fmt.Println(s.Values)
	fmt.Println("===== 出栈 =====")
}