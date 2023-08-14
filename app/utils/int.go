package utils

import (
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
