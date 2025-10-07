package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	checkNext := true

	for _, char := range s {

		if checkNext == true { // if it's the first letter of a word

			if char >= 'a' && char <= 'z' { // check if lowercase
				return false
			}

			checkNext = false
		}

		if char == ' ' {
			checkNext = true
		}
        
	}

	return true
}
