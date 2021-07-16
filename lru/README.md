# LRU (Least Recently Used)

## 实现思路
维护一个队列, 如果某条记录被访问了, 则移动到队尾, 那么队首就是最近最少访问的数据, 淘汰掉队首的记录即可.

## 核心数据结构
使用双向链表 + 哈希表来实现