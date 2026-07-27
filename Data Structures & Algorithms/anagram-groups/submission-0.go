func groupAnagrams(strs []string) [][]string {
	resultAnagrams := [][]string{}
	seen := make([]bool, len(strs))

	// use an index iterate per word
	for i := 0; i < len(strs); i++ {
		if seen[i] {
			continue
		}

		// reference word
		word1 := strs[i]
		seen[i] = true
		// temp anagrams string array
		anagrams := []string{word1}

		// use an index to get the 2nd word
		for j := i+1; j < len(strs); j++ {

			// if len of word1 and word2 do not much
			// break from j-loop and append to resultAnagrams
			if len(strs[i]) != len(strs[j]) {
				continue
			}

			// create frequency map
			counts := make(map[rune]int)

			// initialize freqency map with word1 characters
			for _, c := range word1 {
				counts[c]++
			}

			// decrement when characters are seen in the map
			for _, c := range strs[j] {
				counts[c]--
			}

			// if values are all 0, append to anagrams array
			// if a value is != 0, break from j-loop
			isAnagram := true
			for _, count := range counts {
				if count != 0 {
					isAnagram = false
					break
				} 
			}
			
			// append anagram to anagrams array
			if isAnagram { 
				anagrams = append(anagrams, strs[j])
				seen[j] = true
			}
		}

		// append anagrams array to resultAnagrams
		// move index up
		resultAnagrams = append(resultAnagrams, anagrams)
	}

	return resultAnagrams
}