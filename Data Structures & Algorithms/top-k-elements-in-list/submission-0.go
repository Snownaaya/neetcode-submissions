func topKFrequent(nums []int, k int) []int {
	exists := make(map[int]int)

	for _, value := range nums {
		exists[value]++
	}

	freqList := make([][]int, len(nums)+1)

	for number, freq := range exists {
		freqList[freq] = append(freqList[freq], number)
	}

	result := make([]int, 0)

outer:
	for freq := len(freqList) - 1; freq >= 0; freq-- {
		for _, number := range freqList[freq] {
			result = append(result, number)
			if len(result) == k {
				break outer
			}
		}
	}

	return result
}
