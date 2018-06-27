package bowling

import "fmt"

func Score(bowls []int) int {

	var sc int
	var bowlInFrame int
	var frameNum int

	for i := range bowls {

		frameNum++
		bowlInFrame++

		if frameNum <= 10 {
			sc += bowls[i]
		}

		if bowls[i] == 10 && frameNum <= 10 { // strike
			sc += bowlAt(i+1, bowls) + bowlAt(i+2, bowls)
			bowlInFrame = 0 // reset
			continue
		}

		if bowlInFrame == 2 && bowls[i-1]+bowls[i] == 10 { // spare
		fmt.Println("Spare!")
			sc += bowlAt(i+1, bowls)
			bowlInFrame = 0 // reset
		}
	}

	return sc
}

// bowlAt fetches the bowl score from the index position i, if it exists
func bowlAt(i int, bowls []int) int {
	if len(bowls) > i {
		return bowls[i]
	}
	return 0
}
