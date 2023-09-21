package utils

import "sort"

type FloatSliceType struct{}

var FloatSlice FloatSliceType

type OF struct {
	K int
	V float64
}

func (slc *FloatSliceType) RandIdFromFloatRates(r []float64) int {
	var oS []*OF
	var total, comp float64

	for k, v := range r {
		oS = append(oS, &OF{
			K: k,
			V: v,
		})
		total += v
	}
	sort.SliceStable(oS, func(i, j int) bool {
		return oS[i].V > oS[j].V
	})
	rId := Float.RandRange(0, total)
	for _, o := range oS {
		comp += o.V
		if rId <= comp {
			return o.K
		}
	}
	return -1
}
