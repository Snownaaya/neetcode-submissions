func isAnagram(s string, t string) bool {
	mapS := make(map[rune]int)

	for _, value := range s {
		v, ok := mapS[value]

		if ok {
			mapS[value] = v + 1
		} else {
			mapS[value] = 1
		}
	}

	for _, value := range t {
		v, ok := mapS[value]

		if ok {
			mapS[value] = v - 1
			if v == 1 {
				delete(mapS, value)
			}
		} else {
			return false
		}
	}

	return len(mapS) == 0
}

