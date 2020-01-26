package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	// if list empty, return false
	if len(list) == 0 {
		return false
	}

	// for each number in list
	for _, n := range list {
		// if number is the number, return true
		if n == num {
			return true
		}
	}

	// return false
	return false
}
