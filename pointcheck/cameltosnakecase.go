package piscine

func CamelToSnakeCase(s string) string {

	if s == "" {
		return ""
	}
	// If the string is not camelCase, return the string unchanged.

	for _, char := range s {

		if char >= '0' && char <= '9' {
			return s
		}
	}

	lastuppercase := s[len(s)-1]

	if lastuppercase >= 'A' && lastuppercase <= 'Z' {

		return s
	}

	for i := 1; i < len(s); i++ {

		if s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {

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

		if char >= 'A' && char <= 'Z' {

			result = result + "_" + string(char)
		}else{
            result+=  string(char)
        }
	}
	return result
}
