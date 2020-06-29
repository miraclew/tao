package golang

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"github.com/miraclew/tao/tools/tao/proto"
)

type ServiceMapper struct {
	TypeMapper
}

func (m ServiceMapper) Map(s *proto3.Service) (*Service, error) {
	st, name := proto.ParseServiceName(s.Name)
	if st == proto.ServiceUnknown {
		return nil, fmt.Errorf("unknown service type for name %s", s.Name)
	}
	v := &Service{
		Name:    name,
		Type:    st,
		Methods: []Method{},
	}

	for _, entry := range s.Entry {
		if entry.Method == nil {
			continue
		}

		p, _ := m.TypeMapper.Map(entry.Method.Request)
		r, _ := m.TypeMapper.Map(entry.Method.Response)
		if entry.Method.StreamingRequest {
			p = "chan " + p
		}
		if entry.Method.StreamingResponse {
			r = "chan " + r
		}

		name := entry.Method.Name
		method := Method{
			Name:     name,
			Request:  p,
			Response: r,
		}
		v.Methods = append(v.Methods, method)
	}

	return v, nil
}
