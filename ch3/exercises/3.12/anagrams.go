package anagrams

func anagrams(s1, s2 string) bool {
	j := 0
	for _, c := range s1 {
		if hasIt(s2, c) {
			j++
		}
	}
	return len(s1) == j
}

func hasIt(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
