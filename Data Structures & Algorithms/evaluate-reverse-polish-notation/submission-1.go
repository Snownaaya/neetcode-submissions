func evalRPN(tokens []string) int {
	stack := []int{}

	operations := map[string]func(int, int) int{
		"+": func(number1, number2 int) int {
			return number1 + number2
		},
		"*": func(number1, number2 int) int {
			return number1 * number2
		},
		"-": func(number1, number2 int) int {
			return number1 - number2
		},
		"/": func(number1, number2 int) int {
			return number1 / number2
		},
	}

	for _, token := range tokens {
		calc, ok := operations[token]

		if ok {
			number1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := calc(number2, number1)
			stack = append(stack, result)
		} else {
			n, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}
			stack = append(stack, n)
		}
	}

	return stack[len(stack)-1]
}

