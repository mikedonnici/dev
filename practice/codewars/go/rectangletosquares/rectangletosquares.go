package rectangletosquares

func SquaresInRect(length int, width int) []int {

	if length == width || length*width <= 1 {
		return nil
	}

	var squareSizes []int
	for {
		if length*width == 0 {
			break
		}
		if width > length {
			width, length = length, width
		}
		squareSizes = append(squareSizes, width)
		length = length - width
	}

	return squareSizes
}
