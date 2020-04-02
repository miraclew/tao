package golang

import (
	"strings"
)

type ProtoGolang struct {
	Name     string // api name, cap camel
	Module   string
	URL      string
	Enums    []*Enum
	Messages []*Message
	Service  *Service
	Event    *Service
}

func (p *ProtoGolang) Pkg() string {
	return strings.ToLower(p.Name)
}

// The first model
func (p *ProtoGolang) PrimaryModel() *Message {
	return p.ModelMessages()[0]
}

func (p *ProtoGolang) ModelMessages() []*Message {
	var ms []*Message
	for _, e := range p.Messages {
		if !e.Model {
			continue
		}
		ms = append(ms, e)
	}
	return ms
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
	Model  bool
}

func (m Message) InsertFields() []string {
	var fs []string
	for _, field := range m.Fields {
		if field.Name == "Id" || field.Name == "CreatedAt" || field.Name == "UpdatedAt" {
			continue
		}

		fs = append(fs, field.Name)
	}
	return fs
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
