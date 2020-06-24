package tsql

import (
	"fmt"
	"strconv"
	"strings"
)

type QueryParams struct {
	Ids    []int64
	Filter Filter
	Sort   string
	Offset int64
	Limit  int32
}

type GetParams struct {
	Id     int64
	Filter Filter
}

type Values map[string]interface{}

func (vs Values) UpdateSet() (string, []interface{}) {
	var ss []string
	var values []interface{}
	for k, v := range vs {
		// check increment in update statement, e.g. count=count+1
		if s, ok := v.(string); ok && strings.HasPrefix(s, k) {
			s = strings.TrimSpace(strings.TrimPrefix(s, k))
			if len(s) >= 2 && (s[0:1] == "+" || s[0:1] == "-") {
				ss = append(ss, fmt.Sprintf("%s=%s%s?", k, k, s[0:1]))
				d, _ := strconv.ParseInt(strings.TrimSpace(s[1:]), 10, 64)
				values = append(values, d)
				continue
			}
		}
		ss = append(ss, k+"=?")
		values = append(values, v)
	}
	return strings.Join(ss, ","), values
}

type Filter map[string]interface{}

func (f Filter) Conditions() (string, []interface{}) {
	var cond []string
	var values []interface{}
	for k, v := range f {
		cond = append(cond, k+"=?")
		values = append(values, v)
	}
	return strings.Join(cond, " AND "), values
}
