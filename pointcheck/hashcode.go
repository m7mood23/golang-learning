package piscine

func HashCode(dec string) string {

	result := ""

	for _, char := range dec {

		newhash := (int(char) + (len(dec))) % 127

		if newhash < 32 {

			newhash = newhash + 33
		}

		result = result + string(newhash)
	}

	return result

}
