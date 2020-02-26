package slice

type Uint64Slice []uint64

func (s Uint64Slice) Has(item uint64) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func (s Uint64Slice) Sub(items Uint64Slice) Uint64Slice {
	var res Uint64Slice
	for _, v := range s {
		if !items.Has(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s Uint64Slice) Equals(s2 Uint64Slice) bool {
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
