package piscine

func RetainFirstHalf(str string) string {

	// If the length of the string is odd, round it down.

	if str == "" {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	half := len(str) / 2

	return str[:half]

}
