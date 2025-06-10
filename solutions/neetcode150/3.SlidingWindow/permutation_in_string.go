package main


// leetcode #567
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	l, r, matches := 0, len(s1), 0
	
	s1freq, s2freq := [26]int{}, [26]int{}

	for i := 0; i < len(s1); i++ {
		s1freq[s1[i] - byte('a')]++
		s2freq[s2[i] - byte('a')]++
	}

	for i := 0; i < 26; i++ {
		if s1freq[i] == s2freq[i] {
			matches++
		}
	}

	for r < len(s2) {
		if matches == 26 {
			return true
		}
        
		// add new character from the right
		index := s2[r] - byte('a')
		s2freq[index]++
		if s1freq[index] == s2freq[index]{
			matches++
		} else if s1freq[index] + 1 == s2freq[index] {
			matches--
		}

		// remove character from the left
		index = s2[l] - byte('a')
		s2freq[index]--
		if s1freq[index] == s2freq[index]{
			matches++
		} else if s1freq[index]-1 == s2freq[index] {
			matches--
		}

		l++
		r++
	}

	return matches == 26

}
