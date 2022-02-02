package imageMaker

import "math/rand"

func randRange(min, max int) int {
	// Random number in the range [min, max)
	return rand.Intn(max-min) + min
}
