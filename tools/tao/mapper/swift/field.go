package swift

import (
	"github.com/miraclew/tao/tools/tao/mapper/ir"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"strings"
)

type fieldMapper struct {
	ir.TypeMapper
	Enums []*ir.Enum
}

func (f fieldMapper) Map(field *proto3.Field) (*ir.Field, error) {
	elemType, err := f.TypeMapper.Map(field.Type)
	if err != nil {
		return nil, err
	}
	destType := elemType
	if field.Optional {
		destType += "?"
	}

	isEnum := false
	if field.Type.Reference != "" {
		for _, e := range f.Enums {
			if e.Name == field.Type.Reference {
				isEnum = true
				break
			}
		}
	}

	t := &Type{
		name:     destType,
		scalar:   field.Type.Scalar != proto3.None,
		isMap:    field.Type.Map != nil,
		repeated: field.Repeated,
		enum:     isEnum,
	}

	name := strings.ToLower(field.Name[0:1]) + field.Name[1:]
	v := &ir.Field{
		Name: name,
		Type: t,
	}

	return v, nil
}
