package slice

type UintSlice []uint

func (s UintSlice) Has(item uint) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func (s UintSlice) Sub(items UintSlice) UintSlice {
	var res UintSlice
	for _, v := range s {
		if !items.Has(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s UintSlice) Equals(s2 UintSlice) bool {
	if len(s) != len(s2) {
		return false
	}

	for i, v := range s {
		if v != s2[i] {
			return false
		}
	}
	return true
}
