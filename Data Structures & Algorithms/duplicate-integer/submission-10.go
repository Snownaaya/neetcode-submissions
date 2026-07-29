func hasDuplicate(nums []int) bool {
	newMap := make(map[int]int)

	for index, value := range nums {
		_, ok := newMap[nums[index]]

		if ok {
			return true
		} else {
			newMap[value] = nums[index]
		}
	}

	return false
}