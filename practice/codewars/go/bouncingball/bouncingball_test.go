package bouncingball

import "testing"

func TestBouncingBall(t *testing.T) {

	cases := []struct {
		height       float64
		bounceFactor float64
		windowHeight float64
		views        int
	}{
		{1, 0.66, 1, 1},
	}

	for _, c := range cases {
		got := BouncingBall(c.height, c.bounceFactor, c.windowHeight)
		want := c.views
		if got != want {
			t.Errorf("BouncingBall() = %v, want %v", got, want)
		}
	}

}
