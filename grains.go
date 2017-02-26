package grains

import (
	"errors"
)

const testVersion = 1

var grains [65]uint64

func Square(n int) (grains uint64, err error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}
	return x(n), nil
}

func x(n int) uint64 {
	if grains[n] != 0 {
		return grains[n]
	}
	if n == 1 {
		grains[n] = 1
	} else {
		grains[n] = x(n-1) * 2
	}
	return grains[n]
}

func Total() uint64 {
	var result uint64
	for i := 1; i < 65; i++ {
		if grains[i] == 0 {
			grains[i] = x(i)
		}
		result += grains[i]
	}
	return result
}
