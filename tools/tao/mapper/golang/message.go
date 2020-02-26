package golang

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type MessageMapper struct {
	FieldMapper
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

		if entry.Field == nil {
			continue
		}

		field, err := mm.FieldMapper.Map(entry.Field)
		if err != nil {
			return nil, err
		}
		s.Fields = append(s.Fields, *field)
	}
	return s, nil
}
