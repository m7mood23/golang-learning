package piscine

func StrRev(s string) string {

	result := ""

	for i := len(s) - 1; i >= 0; i-- {

		result = result + string(s[i])
	}

	return result

}
