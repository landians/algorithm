package trie

/*
Trie 前缀树
			root
            / \
          a    b
         /      \
        b        c
       / \        \
      c   d        e
*/
type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// TrieNode 前缀树节点
type TrieNode struct {
	Pass  int         // 记录建立前缀树时该节点被经过的次数
	End   int         // 记录有多少字符串是以该节点结尾的
	Nexts []*TrieNode // 记录一共有多少条路
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		// 这里因为只存储 26 个小写英文字母
		// 如果想存储更多的字符，可以改成hash表存储， map[byte]*Node
		Nexts: make([]*TrieNode, 26),
	}
}

// Insert Trie树的节点的新增
func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	wordBytes := []byte(word)
	node := t.root

	for i := 0; i < len(wordBytes); i++ {
		index := wordBytes[i] - 'a'
		// 无法直接到达路，则新建路
		if node.Nexts[index] == nil {
			node.Nexts[index] = NewTrieNode()
		}
		node = node.Nexts[index]
		node.Pass++
	}
	node.End++
}

// Search 查询 Trie 树中是否存在某个字符串， 返回值 n > 0 说明这个字符串被加入过 n 次，n == 0 则这个字符串没有被添加过
func (t *Trie) Search(word string) int {
	if len(word) == 0 {
		return 0
	}

	wordBytes := []byte(word)
	node := t.root

	for i := 0; i < len(wordBytes); i++ {
		// 决定往哪条路走
		index := wordBytes[i] - 'a'
		// 无法直接到达路，则直接返回0
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index]
	}
	return node.End
}

// Delete 删除 Trie 树中某个字符串
func (t *Trie) Delete(word string) {
	if t.Search(word) == 0 {
		return
	}

	wordBytes := []byte(word)
	node := t.root

	for i := 0; i < len(wordBytes); i++ {
		// 决定往哪条路走
		index := wordBytes[i] - 'a'
		node.Nexts[index].Pass--
		if node.Nexts[index].Pass == 0 {
			node.Nexts[index] = nil
			return
		}
		node = node.Nexts[index]
	}
	node.End--
}

// PrefixNumber 查询 Trie 树中以 prefix 为前缀的字符串的个数
func (t *Trie) PrefixNumber(prefix string) int {
	if len(prefix) == 0 {
		return 0
	}

	prefixBytes := []byte(prefix)
	node := t.root

	for i := 0; i < len(prefixBytes); i++ {
		// 决定往哪条路走
		index := prefixBytes[i] - 'a'
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index]
	}
	return node.Pass
}
