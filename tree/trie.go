package tree

// 字典树的实现

// 字典树中的每个节点的表示
type node struct {
	isWord bool             // 到这个节点是不是单词
	next   map[string]*node // 指向下一个节点
}

type Trie struct {
	root *node // 字典树的根
	size int   // 字典树中单词的个数
}

// 获取trie中单词的个数
func (t *Trie) Init() {
	t.size = 0
	n := node{false, make(map[string]*node)}
	t.root = &n
}
func (t *Trie) GetSize() int {
	return t.size
}

// 向trie中添加一个单词
func (t *Trie) Add(word string) {
	// 遍历字符串, 将字符串中的字符一个一个的插入到trie的节点中
	cur := t.root
	for _, ch := range word {
		if _, ok := cur.next[string(ch)]; !ok { // 如果这个的字符不存在, 就将该字符作为一个node添加到next中
			// 将该字符封装为node
			n := node{false, make(map[string]*node)}
			cur.next[string(ch)] = &n
		}
		cur = cur.next[string(ch)]
	}

	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

// 判断一个单词是否在trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, v := range word {
		if _, ok := cur.next[string(v)]; !ok { // 如果在该节点中有下一个字符
			return false
		}
		cur = cur.next[string(v)]
	}

	return cur.isWord
}

// 判断trie是否有单词以prefix为前缀
func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, v := range prefix {
		if _, ok := cur.next[string(v)]; !ok {
			return false
		}

		cur = cur.next[string(v)]
	}

	return true
}
