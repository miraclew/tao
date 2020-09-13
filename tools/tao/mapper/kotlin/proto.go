package kotlin

import "fmt"

type Type struct {
	name     string
	scalar   bool
	enum     bool
	isMap    bool
	repeated bool
}

func (t Type) Name() string {
	return t.name
}

func (t Type) Scalar() bool {
	return t.scalar
}

func (t Type) Enum() bool {
	return t.enum
}

func (t Type) Map() bool {
	return t.isMap
}

func (t Type) Repeated() bool {
	return t.repeated
}

func (t Type) String() string {
	if t.repeated {
		return fmt.Sprintf("List<%s>", t.name)
	}
	return t.name
}
