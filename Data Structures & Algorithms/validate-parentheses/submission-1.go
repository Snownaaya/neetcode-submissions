
func isValid(s string) bool {
	str := map[rune]rune{'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, c := range s {
		expected, isOpened := str[c]

		if isOpened {
			stack = append(stack, expected)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			if top != c {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
