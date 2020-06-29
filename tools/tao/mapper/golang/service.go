package golang

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"github.com/miraclew/tao/tools/tao/proto"
	"strings"
)

type ServiceMapper struct {
	TypeMapper
}

func (m ServiceMapper) Map(s *proto3.Service) (*Service, error) {
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

		p, _ := m.TypeMapper.Map(entry.Method.Request)
		r, _ := m.TypeMapper.Map(entry.Method.Response)
		if entry.Method.StreamingRequest {
			p = "chan " + p
		}
		if entry.Method.StreamingResponse {
			r = "chan " + r
		}

		// process method name
		name := entry.Method.Name
		if st == proto.ServiceSocket {
			if strings.Index(name, "Recv") != -1 {
				continue
			}
			idx := strings.Index(name, "Send")
			if idx != -1 {
				name = strings.TrimPrefix(name, "Send")
			}
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
