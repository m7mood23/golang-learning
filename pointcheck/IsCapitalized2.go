package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	isfirstLetter := true

	for _, char := range s {

		if isfirstLetter == true { // if it's the first letter of a word

			if char >= 'a' && char <= 'z' { // check if lowercase
				return false
			}

			isfirstLetter = false
		}

		if char == ' ' {
			isfirstLetter = true
		}

	}

	return true
}
