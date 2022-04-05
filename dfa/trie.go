/*
 * @Author: your name
 * @Date: 2022-04-04 14:01:35
 * @LastEditTime: 2022-04-05 19:37:48
 * @LastEditors: Please set LastEditors
 * @Description: 字典树 实现方式分为通过map 或是 直接通过数组
 * @FilePath: /dfa/trie.go
 */
package dfa

type TrieNode struct {
	curValue string             // 当前节点记录的文本
	isWord   bool               // 文本是否结束
	next     map[rune]*TrieNode // 子叶子节点
}

type Trie struct {
	size int       // 节点数量
	root *TrieNode // 根节点
}

func NewTrie() *Trie {
	return &Trie{
		0,
		&TrieNode{"", false, make(map[rune]*TrieNode)},
	}
}

// 非递归添加
func (this *Trie) Add(word string) {
	if len(word) == 0 {
		return
	}

	cur := this.root
	for _, v := range word {
		_, ok := cur.next[v]
		if !ok {
			cur.next[v] = &TrieNode{"", false, map[rune]*TrieNode{}}
		}
		cur = cur.next[v]
		cur.curValue = string(v)
	}
	if !cur.isWord {
		cur.isWord = true
		this.size++
	}
}

func (this *Trie) Contains(word string) (node *TrieNode, index int) {
	if len(word) == 0 {
		return nil, 0
	}

	f := false
	cur := this.root
	for k, v := range word {
		if f {
			index = k
			f = false
		}
		if cur.next[v] == nil {
			return nil, index
		}
		cur = cur.next[v]
		if cur.isWord {
			f = true
		}
	}

	if cur.isWord {
		index = len(word)
	}
	return cur, index
}

func (this *Trie) Prefix(word string) []string {
	node, _ := this.Contains(word)
	if node == nil {
		return nil
	}
	return this.Walk(node)
}

func (this *Trie) Walk(node *TrieNode) (ret []string) {
	if node.isWord {
		ret = append(ret, node.curValue)
	}
	for _, v := range node.next {
		ret = append(ret, this.Walk(v)...)
	}
	return
}

func (this *Trie) Check(values string) (target []string, found []string) {
	start := -1
	end := -1
	tmp := ""

	var nextMap map[rune]*TrieNode
	for i, v := range []rune(values) {
		if nextMap == nil {
			curNode, _ := this.Contains(string(v)) // 当前节点
			if curNode != nil {
				end++
				if end == 0 { // 第一次将两个指针指向同一个
					start = i
				}
				tmp = curNode.curValue
				if !curNode.isWord {
					nextMap = curNode.next
				} else {
					target = append(target, tmp)
					tmp = ""
					found = append(found, string([]rune(values)[start:i+1]))
					start = -1
					end = -1
					nextMap = nil
				}
			} else {
				nextMap = nil
				start = -1
				end = -1
			}
		} else {
			curNode, ok := nextMap[v]
			if ok {
				tmp += curNode.curValue
				if !curNode.isWord {
					nextMap = curNode.next
				} else {
					target = append(target, tmp)
					tmp = ""
					found = append(found, string([]rune(values)[start:i+1]))
					start = -1
					end = -1
					nextMap = nil
				}
			} else {
				nextMap = nil
				start = -1
				end = -1
			}
		}
	}
	return
}
