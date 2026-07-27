func removeElement(nums []int, val int) int {
	result := []int{}

	for _, num := range nums {
		if num != val {
			result = append(result, num)
		}
	}

	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}
	
	return len(result)
}
