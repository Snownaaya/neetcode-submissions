func twoSum(nums []int, target int) []int {
	exists := make(map[int]int)

	for index, value := range nums {
		temp := target - value
		v, ok := exists[temp]

		if ok {
			return []int{v, index}
		}else{
			exists[value] = index
		}
	}
	return nil
}

