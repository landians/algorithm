package lfu

import (
	"container/heap"
	"fmt"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

var onEliminate = func(key string, value Value) {
	fmt.Printf("eliminate key = [%s], value = [%v]\n", key, value)
}

func Test_Heap(t *testing.T) {
	entries := []*entry{
		{
			key: "1",
			value: String("a"),
			frequent: 3,
		},
		{
			key: "2",
			value: String("b"),
			frequent: 2,
		},
		{
			key: "3",
			value: String("c"),
			frequent: 1,
		},
	}

	h := &Heap{}
	*h = append(*h, entries...)


	heap.Init(h)

	fmt.Println("========== 初始堆 ==========")

	for _, hv := range *h {
		fmt.Printf("frequent = [%d]\n", hv.frequent)
	}

	fmt.Println("========== 初始堆 ==========")


	(*h)[0].frequent = 10
	heap.Fix(h, 0)

	fmt.Println("========== 堆修改 ==========")

	for _, hv := range *h {
		fmt.Printf("frequent = [%d]\n", hv.frequent)
	}

	fmt.Println("========== 堆修改 ==========")
}

func TestLFUCache_Get(t *testing.T) {
	c := New(1024, onEliminate)
	c.Add("No.1", String("golang"))
	c.Add("No.2", String("java"))
	c.Add("No.3", String("c++"))
	c.Add("No.4", String("rust"))

	if v, ok := c.Get("No.1"); !ok || string(v.(String)) != "golang"{
		t.Fatal("lru cache hit key = No.1 failed!")
	}

	if _, ok := c.Get("No.unknow"); ok {
		t.Fatal("lru cache miss key = No.unknow failed!")
	}
}

func TestLFUCache_Eliminate(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	maxBytes :=len(k1 + k2 + v1 + v2)

	c := New(int64(maxBytes), onEliminate)
	c.Add(k1, String(v1))
	c.Add(k2, String(v2))

	c.Get(k1)
	c.Get(k2)

	for _, hv := range *c.heap {
		fmt.Printf("frequent = [%d]\n", hv.frequent)
	}

	c.Add(k3, String(v3))

	if _, ok := c.Get(k3); ok || c.Len() != 2 {
		t.Fatal("eliminate k1 failed!")
	}
}