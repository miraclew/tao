package golang

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"strings"
)

type ServiceMapper struct {
	TypeMapper
}

func (s2 ServiceMapper) Map(s *proto3.Service) (*Service, error) {
	v := &Service{
		Name:    s.Name,
		Methods: []Method{},
	}

	for _, entry := range s.Entry {
		if entry.Method == nil {
			continue
		}

		p, _ := s2.TypeMapper.Map(entry.Method.Request)
		r, _ := s2.TypeMapper.Map(entry.Method.Response)
		if entry.Method.StreamingRequest {
			p = "chan " + p
		}
		if entry.Method.StreamingResponse {
			r = "chan " + r
		}

		name := entry.Method.Name
		if s.Name == "Event" {
			name = strings.TrimPrefix(name, "Handle")
		}
		method := Method{
			Name:     name,
			Request:  p,
			Response: r,
		}
		v.Methods = append(v.Methods, method)
	}
	return v, nil
}
