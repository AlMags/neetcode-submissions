func isAnagram(s string, t string) bool {
	word1 := make(map[rune]int)
	word2 := make(map[rune]int)

	for _, ch := range s {
		word1[ch]++
	}

	for _, ch := range t {
		word2[ch]++
	}

	if len(word1) != len(word2) {
		return false
	}

	for char, count := range word1 {
		if word2[char] != count {
			return false
		}
	}

	return true
}
