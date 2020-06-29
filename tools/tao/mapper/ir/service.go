package ir

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"github.com/miraclew/tao/tools/tao/proto"
	"strings"
)

type serviceMapper struct {
	TypeMapper
}

func NewServiceMapper(tm TypeMapper) ServiceMapper {
	return &serviceMapper{tm}
}

func (s2 serviceMapper) Map(s *proto3.Service) (*Service, error) {
	st, _ := proto.ParseServiceName(s.Name)
	if st == proto.ServiceUnknown {
		return nil, fmt.Errorf("unknown service type for name %s", s.Name)
	}

	v := &Service{
		Name:    s.Name,
		Type:    st,
		Methods: []Method{},
	}

	for _, entry := range s.Entry {
		if entry.Method == nil {
			continue
		}
		p, _ := s2.TypeMapper.Map(entry.Method.Request)
		r, _ := s2.TypeMapper.Map(entry.Method.Response)

		name := strings.ToLower(entry.Method.Name[0:1]) + entry.Method.Name[1:]
		method := Method{
			Name:     name,
			Request:  p,
			Response: r,
		}
		v.Methods = append(v.Methods, method)
	}
	return v, nil
}
