package dart

import (
	"github.com/miraclew/tao/pkg/slice"
	"github.com/miraclew/tao/tools/tao/mapper/ir"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"strings"
)

type protoMapper struct {
}

func NewProtoMapper() ir.ProtoMapper {
	return &protoMapper{}
}

func (p protoMapper) Map(proto *proto3.Proto) (*ir.ProtoIR, error) {
	var tm typeMapper
	em := ir.NewEnumMapper()

	// enums
	var enums []*ir.Enum
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

	fm := &fieldMapper{
		TypeMapper: tm,
		Enums:      enums,
	}

	mm := ir.NewMessageMapper(fm)
	sm := ir.NewServiceMapper(tm)

	var ignoreMessages = slice.StringSlice{"Time", "Any", "Empty", "Key"}
	var messages []*ir.Message
	for _, entry := range proto.Entries {
		if entry.Message != nil {
			if ignoreMessages.Has(entry.Message.Name) {
				continue
			}
			if strings.HasSuffix(entry.Message.Name, "Event") {
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
	var service *ir.Service
	var event *ir.Service
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

	resource, err := ir.FileOption(proto, "resource")
	if err != nil {
		return nil, err
	}
	app, err := ir.FileOption(proto, "app")
	if err != nil {
		return nil, err
	}

	protoIR := &ir.ProtoIR{
		App:      app,
		Name:     resource,
		Enums:    enums,
		Service:  service,
		Event:    event,
		Messages: messages,
	}

	return protoIR, nil
}
