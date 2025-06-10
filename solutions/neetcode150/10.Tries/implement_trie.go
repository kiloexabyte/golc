package tries

// Leetcode #208
type Trie struct {
    root TrieNode
}

type TrieNode struct {
	end bool
	value rune
	children map[rune]*TrieNode
}

func NewTrieNode1() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func NewTrieNode2(value rune, end bool) *TrieNode {
	return &TrieNode{end, value, make(map[rune]*TrieNode)}
}


func Constructor() Trie {
	rootNode := NewTrieNode1()
	return Trie{*rootNode}
}


func (this *Trie) Insert(word string)  {
	curr := &this.root

	for _, c := range word {
		_, ok := curr.children[c]
		if !ok {
			curr.children[c] = NewTrieNode2(c, false)
		}

		curr = curr.children[c]
	}

	curr.end = true
}


func (this *Trie) Search(word string) bool {
	curr := &this.root

	for _, c := range word {
		_, ok := curr.children[c]
		if !ok {
			return false
		}

		curr = curr.children[c]
	}

	return curr.end
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := &this.root

	for _, c := range prefix {
		_, ok := curr.children[c]
		if !ok {
			return false
		}

		curr = curr.children[c]
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */