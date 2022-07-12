package multiple_strings_sort

type MSort struct {
	s []string
}

func NewMSort(s []string) *MSort {
	return &MSort{
		s: s,
	}
}

func (m *MSort) Len() int {
	return len(m.s)
}

func (m *MSort) Less(i, j int) bool {
	li := len(m.s[i])
	lj := len(m.s[j])

	mlen := li
	if li > lj {
		mlen = lj
	}
	res := true
	for k := 0; k < mlen; k++ {
		if m.s[i][k] < m.s[j][k] {
			return true
		} else if m.s[i][k] > m.s[j][k] {
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

func (m *MSort) Swap(i, j int) {
	if ok := m.Less(i, j); ok {
		m.s[i], m.s[j] = m.s[j], m.s[i]
	}
}
