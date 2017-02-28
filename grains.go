package grains

import (
	"errors"
	"math"
)

const testVersion = 1

var grains [65]uint64

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}
	if grains[n] != 0 {
		return grains[n], nil
	}
	grains[n] = uint64(math.Pow(2, float64(n)-1))
	return grains[n], nil
}

func Total() uint64 {
	var result uint64
	for i := 1; i < 65; i++ {
		if grains[i] == 0 {
			grains[i] = uint64(math.Pow(2, float64(i)-1))
		}
		result += grains[i]
	}
	return result
}
