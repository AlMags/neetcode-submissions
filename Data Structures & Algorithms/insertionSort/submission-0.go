// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	result := [][]Pair{}

	for i, _ := range pairs {
		j := i - 1

		for j >= 0 && pairs[j + 1].Key < pairs[j].Key {
			pairs[j + 1], pairs[j] = pairs[j], pairs[j + 1]
			j--
		}

		dup := make([]Pair, len(pairs))
		copy(dup, pairs)
		result = append(result, dup)
	}

	return result
}
