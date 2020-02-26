package ir

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type ProtoMapper interface {
	Map(proto *proto3.Proto) (*ProtoIR, error)
}

type TypeMapper interface {
	Map(t *proto3.Type) (string, error)
}

type EnumMapper interface {
	Map(e *proto3.Enum) (*Enum, error)
}

type FieldMapper interface {
	Map(f *proto3.Field) (*Field, error)
}

type MessageMapper interface {
	Map(m *proto3.Message) (*Message, error)
}

type ServiceMapper interface {
	Map(s *proto3.Service) (*Service, error)
}
