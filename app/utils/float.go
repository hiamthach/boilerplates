package utils

import (
	"math/rand"
	"time"
)

type floatType struct{}

var Float floatType

func (i *floatType) RandRange(l, u float64) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Float64()*(u-l) + l
}
