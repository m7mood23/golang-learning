package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	isFirstLetter := true

	for _, char := range s {

		if isFirstLetter == true { // if it's the first letter of a word

			if char >= 'a' && char <= 'z' { // check if lowercase
				return false
			}

			isFirstLetter = false
		}

		if char == ' ' {
			isFirstLetter = true
		}

	}

	return true
}
