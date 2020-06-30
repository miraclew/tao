{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"{{.Module}}/{{.Pkg}}"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Repo interface {
	Query(ctx context.Context, q *QueryParams) ([]*{{.Pkg}}.{{.Name|title}}, error)
	Get(ctx context.Context, q *GetParams) (*{{.Pkg}}.{{.Name|title}}, error)
	Update(ctx context.Context, v Values, id int64) error
	Create(ctx context.Context, req *{{.Pkg}}.{{.Name|title}}) (int64, error)
	Delete(ctx context.Context, id int64) error
}

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
