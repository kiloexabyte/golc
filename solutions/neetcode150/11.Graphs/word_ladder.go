package graphs

func ladderLength(beginWord string, endWord string, wordList []string) int {
    found := false
    for _, word := range wordList {
        if word == endWord {
            found = true
            break
        }
    }
    if !found {
        return 0
    }

	nb := make(map[string][]string)
	nb[beginWord] = []string{}

	for _, word := range wordList {
		for j := range len(word) {
			pattern := word[:j] + "*" + word[j+1:]
			nb[pattern] = append(nb[pattern], word)
		}
	}

	visit := make(map[string]struct{})
	visit[beginWord] = struct{}{}

	q := []string{beginWord}
	res := 1

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			// pop q
			word := q[0]      // get the first element
			q = q[1:]         // remove the first element

			if word == endWord {
				return res
			}

			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j + 1:]
				for _, neiWord := range nb[pattern] {
					if _, exist := visit[neiWord]; !exist {
						q = append(q, neiWord)
						visit[neiWord] = struct{}{}
					}
				}
			}
		}

		res++
	}

	return 0
}