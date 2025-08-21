package piscine

func RectPerimeter(w, h int) int {
	//Edge Case
	//check if width OR height is negative
	if w < 0 || h < 0 {

		return -1
	}

	// calculate perimeter : 2 * (width + height )
	return 2 * (w + h)

}
