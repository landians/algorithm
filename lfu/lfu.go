package lfu

import (
	"container/heap"
)

// lfu cache 对象
type LFUCache struct {
	maxBytes  int64                         // 允许使用的最大内存字节数
	nBytes    int64                         // 当前已经使用的内存字节数
	heap      *Heap                          // 最小堆
	cache     map[string]*entry             // 哈希表
	OnEvicted func(key string, value Value) // 当一个 entry 被淘汰时触发的回调函数
}

type Heap []*entry

// key
type entry struct {
	key      string
	value    Value
	frequent int64
}

// value
type Value interface {
	Len() int
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].frequent < (*h)[j].frequent
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(*entry))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

func New(maxBytes int64, OnEvicted func(key string, value Value)) *LFUCache {
	c := &LFUCache{
		maxBytes:  maxBytes,
		heap:      &Heap{},
		cache:     make(map[string]*entry),
		OnEvicted: OnEvicted,
	}
	heap.Init(c.heap)
	return c
}

// 根据 key 来从缓存中获取元素
func (c *LFUCache) Get(key string) (value Value, ok bool) {
	// 命中缓存，则对应 key 的访问次数+1, 并调整小根堆
	if e, ok := c.cache[key]; ok {
		e.frequent++
		heap.Fix(c.heap, 0)
		return e.value, true
	}
	return nil, false
}

// 添加/更新 缓存的 key-value 对, 命中缓存，则更新，否则新增
func (c *LFUCache) Add(key string, value Value) {
	if e, ok := c.cache[key]; ok {
		e.frequent++
		heap.Fix(c.heap, 0)
		c.nBytes += int64(value.Len()) - int64(e.value.Len())
		e.value = value
	} else {
		v := &entry{
			key:   key,
			value: value,
		}
		c.heap.Push(v)
		c.cache[key] = v
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.doEliminate()
	}
}

// 淘汰缓存中的元素
func (c *LFUCache) doEliminate() {
	if v, ok := c.heap.Pop().(*entry); ok {
		delete(c.cache, v.key)
		c.nBytes -= int64(len(v.key)) + int64(v.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(v.key, v.value)
		}
	}
}

func (c *LFUCache) Len() int {
	return c.heap.Len()
}

func (c *LFUCache) MaxBytes() int64 {
	return c.maxBytes
}

func (c *LFUCache) CurrentBytes() int64 {
	return c.nBytes
}
