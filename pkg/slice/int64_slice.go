package slice

type Int64Slice []int64

func (s Int64Slice) InterfaceSlice() []interface{} {
	var ss []interface{}
	for _, v := range s {
		ss = append(ss, v)
	}
	return ss
}

func (s Int64Slice) Has(item int64) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}
