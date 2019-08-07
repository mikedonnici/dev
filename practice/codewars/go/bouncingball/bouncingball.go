package bouncingball

// BouncingBall returns the number of times a ball will be visible
// from a window (height) when dropped from height h, and bouncing back to
// a height of bounce * h.
func BouncingBall(h, bounce, window float64) int {

	if !paramsOk(h,bounce,window) {
		return -1
	}

	var views int
	for {
		if h < window {
			break
		}
		views++ // on the way down
		h = bounce * h
		if h > window {
			views++ // on the way up
		}
	}

	return views
}

func paramsOk(h, b, w float64) bool {
	// height cannot be 0, or negative
	if h <= 0 {
		return false
	}
	// window height must be below initial height of ball
	if w > h {
		return false
	}
	// bounce factor must be > 0 and < 1
	if b <= 0 || b >= 1 {
		return false
	}
	return true
}
