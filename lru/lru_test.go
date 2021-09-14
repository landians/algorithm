package lru

import (
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

func TestLRUCache_Get(t *testing.T) {
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

func TestLRUCache_Eliminate(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	maxBytes :=len(k1 + k2 + v1 + v2)

	c := New(int64(maxBytes), onEliminate)
	c.Add(k1, String(v1))
	c.Add(k2, String(v2))

	t.Logf("lru cache's current bytes = [%d]\n", c.CurrentBytes())

	c.Add(k3, String(v3))

	t.Logf("lru cache's current bytes = [%d]\n", c.CurrentBytes())

	if _, ok := c.Get(k1); ok || c.Len() != 2 {
		t.Fatal("eliminate k1 failed!")
	}
}

func TestLRUCache_Remove(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"

	c := New(0, onEliminate)
	c.Add(k1, String(v1))
	c.Add(k2, String(v2))
	c.Add(k3, String(v3))

	t.Logf("lru cache's current bytes = [%d]\n", c.CurrentBytes())

	t.Logf("removed entry = [%v]\n", c.Remove(k1))

	t.Logf("lru cache's current bytes = [%d]\n", c.CurrentBytes())
}