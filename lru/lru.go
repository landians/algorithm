package lru

import "container/list"

// lru cache 对象
type LRUCache struct {
	maxBytes  int64                         // 允许使用的最大内存字节数
	nBytes    int64                         // 当前已经使用的内存字节数
	ll        *list.List                    // 双向链表
	cache     map[string]*list.Element      // 哈希表
	OnEvicted func(key string, value Value) // 当一个 entry 被淘汰时触发的回调函数
}

// key
type entry struct {
	key   string
	value Value
}

// value
type Value interface {
	Len() int
}

func New(maxBytes int64, OnEvicted func(key string, value Value)) *LRUCache {
	return &LRUCache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

// 根据 key 来从缓存中获取元素
func (c *LRUCache) Get(key string) (value Value, ok bool) {
	// 命中缓存, 将对应元素移到队尾
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		v := ele.Value.(*entry)
		return v.value, true
	}
	return nil, false
}

// 添加/更新 缓存的 key-value 对, 命中缓存，则更新，否则新增
func (c *LRUCache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		v := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(v.value.Len())
		v.value = value
	} else {
		v := &entry{
			key:   key,
			value: value,
		}
		ele := c.ll.PushFront(v)
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.doEliminate()
	}
}

// 删除缓存中的元素
func (c *LRUCache) Remove(key string) *entry {
	if ele, ok := c.cache[key]; ok {
		c.ll.Remove(ele)
		v := ele.Value.(*entry)
		delete(c.cache, v.key)
		c.nBytes -= int64(len(v.key)) + int64(v.value.Len())
		return v
	}
	return nil
}

// 淘汰缓存中的元素
func (c *LRUCache) doEliminate() {
	ele := c.ll.Back()
	if ele != nil {
		v := ele.Value.(*entry)
		c.Remove(v.key)
		if c.OnEvicted != nil {
			c.OnEvicted(v.key, v.value)
		}
	}
}

func (c *LRUCache) Len() int {
	return c.ll.Len()
}

func (c *LRUCache) MaxBytes() int64 {
	return c.maxBytes
}

func (c *LRUCache) CurrentBytes() int64 {
	return c.nBytes
}
