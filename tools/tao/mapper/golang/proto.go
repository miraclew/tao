package golang

import "strings"

type ProtoGolang struct {
	Name     string // api name, cap camel
	Module   string
	URL      string
	Enums    []*Enum
	Messages []*Message
	Service  *Service
	Event    *Service
}

func (p ProtoGolang) Pkg() string {
	return strings.ToLower(p.Name)
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
	Model  bool
}

type Field struct {
	Name string
	Type Type
	Tags string
}

type Type struct {
	Name     string
	Scalar   bool
	Enum     bool
	Map      bool
	Repeated bool
}

func (t Type) String() string {
	var s = t.Name
	if !t.Scalar && !t.Enum && !t.Map && t.Name != "time.Time" {
		s = "*" + s
	}
	if t.Repeated {
		return "[]" + s
	}
	return s
}
