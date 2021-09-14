package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := New(10)

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	fmt.Println("===== 入队 =====")
	fmt.Println(q.Values)
	fmt.Println("===== 入队 =====")

	fmt.Println()

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	fmt.Println("===== 出队 =====")
	fmt.Println(q.Values)
	fmt.Println("===== 出队 =====")
}
