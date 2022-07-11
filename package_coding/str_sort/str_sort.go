package strsort

type StrSort struct {
	sMap []string
}

func NewStrSort(m []string) *StrSort {
	return &StrSort{
		sMap: m,
	}
}

func (s *StrSort) Len() int {
	return len(s.sMap)
}

func (s *StrSort) Less(i, j int) bool {
	li := len(s.sMap[i])
	lj := len(s.sMap[j])

	mlen := li
	if li > lj {
		mlen = lj
	}
	res := true
	for k := 0; k < mlen; k++ {
		if s.sMap[i][k] < s.sMap[j][k] {
			return true
		} else if s.sMap[i][k] > s.sMap[j][k] {
			return false
		}
	}
	// 如果前mlen都相等
	if res {
		if li < lj {
			return true
		}
	}

	return false
}

func (s *StrSort) Swap(i, j int) {
	if ok := s.Less(i, j); ok {
		s.sMap[i], s.sMap[j] = s.sMap[j], s.sMap[i]
	}
}
