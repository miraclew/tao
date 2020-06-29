package ir

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/proto"
)

// ProtoIR is proto intermediate representative
type ProtoIR struct {
	Name     string // api name, cap camel
	App      string
	Enums    []*Enum
	Services []*Service
	Messages []*Message
}

type Service struct {
	Name    string
	Type    proto.ServiceType
	Methods []Method
}

type Method struct {
	Name     string
	Request  string
	Response string
}

type Enum struct {
	Message string // which message defines this enum, empty if enums defined in file
	Name    string
	Values  []Value
}

type Value struct {
	Name  string
	Text  string // text of the value, use name if empty
	Value int
}

func (v Value) String() string {
	if v.Text != "" {
		return v.Text
	}
	return v.Name
}

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type Type
	Tags string
}

type Type interface {
	Name() string
	Scalar() bool
	Enum() bool
	Map() bool
	Repeated() bool
	fmt.Stringer
}
