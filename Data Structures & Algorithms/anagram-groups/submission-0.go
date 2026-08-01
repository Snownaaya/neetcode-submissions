func groupAnagrams(strs []string) [][]string {
	count := make(map[[26]int][]string)

	for _, value := range strs {
		key := [26]int{}

		for _, char := range value {
			key[char - 'a'] += 1
		}

		count[key] = append(count[key], value)
	}

	matrix := make([][]string, 0)

	for _, value := range count {
		matrix = append(matrix, value)
	}

	return matrix
}
