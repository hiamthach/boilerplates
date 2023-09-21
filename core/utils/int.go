package utils

import (
	"math"
	"math/rand"
	"time"
)

type intType struct{}

var Int intType

func (i *intType) RandRange(l, u int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(u-l+1) + l
}

func (i *intType) Min(a, b int64) int64 {
	min := math.Min(float64(a), float64(b))
	return int64(min)
}

func (i *intType) Max(a, b int64) int64 {
	min := math.Max(float64(a), float64(b))
	return int64(min)
}
