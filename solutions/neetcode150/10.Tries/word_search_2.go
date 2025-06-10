package tries

type TrieNode struct {
	children map[byte]*TrieNode
	end      bool
	value    byte
}

func NewTrieNode(val byte) *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		end:      false,
		value:    val,
	}
}

func (t *TrieNode) AddWord(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.children[c]; !ok {
			cur.children[c] = NewTrieNode(c)
		}
		cur = cur.children[c]
	}
	cur.end = true
}

// Leetcode #212
func findWords(board [][]byte, words []string) []string {
	root := NewTrieNode(' ')
	for _, w := range words {
		root.AddWord(w)
	}

	rows := len(board)
	if rows == 0 {
		return nil
	}
	cols := len(board[0])
	res := make(map[string]struct{})
	visit := make(map[[2]int]struct{})

	var dfs func(r, c int, node *TrieNode, word []byte)
	dfs = func(r, c int, node *TrieNode, word []byte) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		if _, ok := visit[[2]int{r, c}]; ok {
			return
		}
		ch := board[r][c]
		next, ok := node.children[ch]
		if !ok {
			return
		}

		visit[[2]int{r, c}] = struct{}{}
		word = append(word, ch)
		if next.end {
			res[string(word)] = struct{}{}
		}

		dfs(r-1, c, next, word)
		dfs(r+1, c, next, word)
		dfs(r, c-1, next, word)
		dfs(r, c+1, next, word)

		delete(visit, [2]int{r, c})
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, root, []byte{})
		}
	}

	wordsFound := make([]string, 0, len(res))
	for w := range res {
		wordsFound = append(wordsFound, w)
	}
	return wordsFound
}
