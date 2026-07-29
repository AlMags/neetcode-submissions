func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)

	for i, num := range nums {
		difference := target - num
		if j, seen := prevMap[difference]; seen {
			return []int{j, i}
		}
		prevMap[num] = i
	}

	return []int{}
}
