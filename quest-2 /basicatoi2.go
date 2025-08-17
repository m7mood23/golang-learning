package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, value := range s {
		if value >= '0' && value <= '9' {
			result = result*10 + int(value-'0')
		} else {
			return 0
		}
	}
	return result
}
