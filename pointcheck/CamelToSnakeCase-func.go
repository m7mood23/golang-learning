package piscine

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func CamelToSnakeCase(s string) string {

	if s == "" { //keep it
		return ""
	}
	// If the string is not camelCase, return the string unchanged.

	for _, char := range s {

		if isNumber(char) == true {
			return s
		}
	}

	lastuppercase := s[len(s)-1]

	if isUpper(rune(lastuppercase)) == true {
		return s
	}

	for i := 1; i < len(s); i++ {

		if isUpper(rune(s[i])) == true && isUpper(rune(s[i-1])) == true {
			return s
		}
	}

	// If the string is camelCase, return the snake_case version of the string.

	result := ""
	for index, char := range s {

		if index == 0 {
			result = result + string(char)
			continue
		}

		if isUpper(char) == true {

			result = result + "_" + string(char)
		} else {
			result += string(char)
		}
	}
	return result
}
