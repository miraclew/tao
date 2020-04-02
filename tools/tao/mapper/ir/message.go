package ir

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type messageMapper struct {
	FieldMapper
}

func NewMessageMapper(fm FieldMapper) MessageMapper {
	return &messageMapper{fm}
}

func (mm messageMapper) Map(message *proto3.Message) (*Message, error) {
	s := &Message{
		Name:   message.Name,
		Fields: []Field{},
	}

	for _, entry := range message.Entries {
		if entry.Field == nil || entry.Field.Type.Reference == "Key" {
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
