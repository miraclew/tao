package golang

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type MessageMapper struct {
	FieldMapper
	UseSnackCase bool
}

func (mm MessageMapper) Map(message *proto3.Message) (*Message, error) {
	s := &Message{
		Name:   message.Name,
		Fields: []Field{},
	}

	for _, entry := range message.Entries {
		if entry.Option != nil {
			if entry.Option.Name == "model" {
				s.Model = true
			}
			continue
		}

		if entry.Field == nil || entry.Field.Type.Reference == "Key" {
			continue
		}

		field, err := mm.FieldMapper.Map(entry.Field)
		if err != nil {
			return nil, err
		}

		tags := ""
		if mm.UseSnackCase {
			tags = fmt.Sprintf(`json:"%s"`, strcase.ToSnake(field.Name))
		}
		if s.Model {
			tags += fmt.Sprintf(` db:"%s"`, field.Name)
		}
		field.Tags = tags
		s.Fields = append(s.Fields, *field)
	}
	return s, nil
}
