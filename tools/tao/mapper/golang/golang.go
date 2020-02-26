package golang

import (
	"fmt"

	"github.com/miraclew/tao/pkg/slice"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

func Map(proto *proto3.Proto) (*ProtoGolang, error) {
	var tm TypeMapper
	em := EnumMapper{}

	// enums
	var enums []*Enum
	for _, entry := range proto.Entries {
		if entry.Message != nil {
			for _, entry := range entry.Message.Entries {
				if entry.Enum != nil {
					e, err := em.Map(entry.Enum)
					if err != nil {
						return nil, err
					}
					e.Message = entry.Message.Name
					enums = append(enums, e)
				}
			}
		} else if entry.Enum != nil {
			e, err := em.Map(entry.Enum)
			if err != nil {
				return nil, err
			}
			enums = append(enums, e)
		}
	}
	mm := MessageMapper{FieldMapper{tm, enums}}
	sm := ServiceMapper{tm}

	var ignoreMessages = slice.StringSlice{"Time", "Any"}
	var messages []*Message
	for _, entry := range proto.Entries {
		if entry.Message != nil {
			if ignoreMessages.Has(entry.Message.Name) {
				continue
			}
			m, err := mm.Map(entry.Message)
			if err != nil {
				return nil, err
			}
			messages = append(messages, m)
		}
	}

	// services
	var service *Service
	var event *Service
	var err error
	for _, entry := range proto.Entries {
		if entry.Service != nil {
			if entry.Service.Name == "Service" {
				service, err = sm.Map(entry.Service)
			} else if entry.Service.Name == "Event" {
				event, err = sm.Map(entry.Service)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	resource, err := FileOption(proto, "resource")
	if err != nil {
		return nil, err
	}

	protoIR := &ProtoGolang{
		Name:     resource,
		Enums:    enums,
		Service:  service,
		Event:    event,
		Messages: messages,
	}

	return protoIR, nil
}

func FileOption(proto *proto3.Proto, option string) (string, error) {
	for _, entry := range proto.Entries {
		if entry.Option != nil && entry.Option.Name == option {
			return *entry.Option.Value.String, nil
		}
	}
	return "", fmt.Errorf("option %s undefined", option)
}
