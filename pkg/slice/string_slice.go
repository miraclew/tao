package slice

type StringSlice []string

func (s StringSlice) Has(item string) bool {
	if s == nil {
		return false
	}
	for _, v := range s {
		if v == item {
			return true
		}
	}

	return false
}

func (s StringSlice) Sub(items StringSlice) StringSlice {
	var res StringSlice
	for _, v := range s {
		if !items.Has(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s StringSlice) Equals(s2 StringSlice) bool {
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

func StringsContains(haystack []string, noodle string) bool {
	if haystack == nil {
		return false
	}
	for _, v := range haystack {
		if v == noodle {
			return true
		}
	}

	return false
}

func UniqueStrings(ss []string, stripEmpty bool) []string {
	uniq := make(map[string]bool)
	for _, v := range ss {
		uniq[v] = true
	}

	var res []string
	for key := range uniq {
		if stripEmpty && key == "" {
			continue
		}
		res = append(res, key)
	}

	return res
}

func StripEmptyString(ss []string) []string {
	res := []string{} // avoid nil slice
	for _, v := range ss {
		if v == "" {
			continue
		}
		res = append(res, v)
	}
	return res
}
