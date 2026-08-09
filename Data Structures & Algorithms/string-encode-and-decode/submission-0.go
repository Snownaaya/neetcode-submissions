type Solution struct {
	result string
	slice  []string
}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for i := 0; i < len(strs); i++ {
		lengthStr := strconv.Itoa(len(strs[i]))

		_, err := builder.WriteString(lengthStr)
		if err != nil {
			return ""
		}

		var r rune = '#'
		b := byte(r)
		errorr := builder.WriteByte(b)
		if errorr != nil {
			return ""
		}

		builder.WriteString(strs[i])
	}

	s.result = builder.String()
	return s.result
}

func (s *Solution) Decode(encoded string) []string {
	for i := 0; i < len(encoded); {
		if i < len(encoded) {
			hasPos := strings.IndexByte(encoded[i:], '#')
			lenthStr := encoded[i : i+hasPos]

			number, err := strconv.Atoi(lenthStr)
			if err != nil {
				return []string{}
			}

			contentStart := i + hasPos + 1
			contentEnd := contentStart + number
			content := encoded[contentStart:contentEnd]

			s.slice = append(s.slice, content)

			i = contentEnd
		}
	}

	return s.slice
}