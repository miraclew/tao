package golang

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
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
		p, _ := s2.TypeMapper.Map(entry.Method.Request)
		r, _ := s2.TypeMapper.Map(entry.Method.Response)

		method := Method{
			Name:     entry.Method.Name,
			Request:  p,
			Response: r,
		}
		v.Methods = append(v.Methods, method)
	}
	return v, nil
}
