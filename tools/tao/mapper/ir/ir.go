package ir

import "fmt"

// ProtoIR is proto intermediate representative
type ProtoIR struct {
	Name     string // api name, cap camel
	URL      string
	Enums    []*Enum
	Service  *Service
	Event    *Service
	Messages []*Message
}

type Service struct {
	Name    string
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
	Value int
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
