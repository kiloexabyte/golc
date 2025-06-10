package tries

// Leetcode #211
type TrieNode struct {
	end      bool
	value    rune
	children map[rune]*TrieNode
}

func TrieNodeConstructor(value rune, isEnd bool) *TrieNode {
	return &TrieNode{
		end:      isEnd,
		value:    value,
		children: make(map[rune]*TrieNode),
	}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: TrieNodeConstructor(' ', false),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i, c := range word {
		endOfWord := i == len(word)-1
		if _, exists := curr.children[c]; !exists {
			curr.children[c] = TrieNodeConstructor(c, endOfWord)
		}
		curr = curr.children[c]
		if endOfWord {
			curr.end = true
		}
	}
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(word, 0, this.root)
}

func dfs(word string, index int, node *TrieNode) bool {
	curr := node
	for i := index; i < len(word); i++ {
		c := rune(word[i])
		if c == '.' {
			for _, child := range curr.children {
				if dfs(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			if _, exists := curr.children[c]; !exists {
				return false
			}
			curr = curr.children[c]
		}
	}
	return curr.end
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
