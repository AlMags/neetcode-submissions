func hasDuplicate(nums []int) bool {
    var items []int
	for _, num := range nums {
		for _, item := range items {
			if num == item {
				return true
			}
		}
		items = append(items, num)
	}
	return false
}
