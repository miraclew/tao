package golang

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type FieldMapper struct {
	TypeMapper
	Enums []*Enum
}

func (f FieldMapper) Map(field *proto3.Field) (*Field, error) {
	elemType, err := f.TypeMapper.Map(field.Type)
	if err != nil {
		return nil, err
	}
	destType := elemType

	isEnum := false
	if field.Type.Reference != "" {
		for _, e := range f.Enums {
			if e.Name == field.Type.Reference {
				isEnum = true
				break
			}
		}
	}

	t := Type{
		Name:     destType,
		Scalar:   field.Type.Scalar != proto3.None,
		Map:      field.Type.Map != nil,
		Repeated: field.Repeated,
		Enum:     isEnum,
	}
	v := &Field{
		Name: field.Name,
		Type: t,
	}

	return v, nil
}
