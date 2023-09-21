package utils

type StrSliceUtils struct{}

var StrSlice StrSliceUtils

func (s *StrSliceUtils) GetSetSize(ss []string) int {
	m := make(map[string]bool)
	for _, v := range ss {
		m[v] = true
	}
	return len(m)
}
func (s *StrSliceUtils) RemoveEmpty(ss []string) (ss2 []string) {
	for _, v := range ss {
		if v == "" {
			continue
		}
		ss2 = append(ss2, v)
	}
	return
}
func (s *StrSliceUtils) Contain(ss []string, sv string) (exist bool) {
	exist = false
	for _, v := range ss {
		if v == sv {
			exist = true
			return
		}
	}
	return
}
func (s *StrSliceUtils) ToMap(ss []string) (m map[string]bool) {
	for _, si := range ss {
		m[si] = true
	}
	return
}
