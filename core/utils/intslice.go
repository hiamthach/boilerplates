package utils

import (
	"sort"
)

type intSlice struct{}

var IntSlice intSlice

type O struct {
	K, V int
}

func (slc *intSlice) RandIdFromRates(r []int) int {
	var (
		oS          []*O
		total, comp int
	)
	for k, v := range r {
		oS = append(oS, &O{
			K: k,
			V: v,
		})
		total += v
	}
	sort.SliceStable(oS, func(i, j int) bool {
		return oS[i].V > oS[j].V
	})
	rId := Int.RandRange(1, total)
	for _, o := range oS {
		comp += o.V
		if rId <= comp {
			return o.K
		}
	}
	return -1
}
