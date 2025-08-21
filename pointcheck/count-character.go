package piscine

func CountChar(str string, c rune) int {

	if str == "" {
		return 0
	}

	count := 0

	for _, char := range str {

		if char == c {
			count++
		}
	}
	return count

}
