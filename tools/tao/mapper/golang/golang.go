package golang

import (
	"fmt"
	"strings"

	"github.com/miraclew/tao/pkg/slice"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

func Map(proto *proto3.Proto, useSnackCase bool) (*ProtoGolang, error) {
	var tm TypeMapper

	// enums
	em := EnumMapper{}
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
	// messages
	mm := MessageMapper{
		FieldMapper:  FieldMapper{tm, enums},
		UseSnackCase: useSnackCase,
	}
	var ignoreMessages = slice.StringSlice{"Time", "Any", "Key"}
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
	sm := ServiceMapper{tm}
	var services []*Service
	for _, entry := range proto.Entries {
		if entry.Service != nil {
			service, err := sm.Map(entry.Service)
			if err != nil {
				return nil, err
			}
			services = append(services, service)
		}
	}

	goPackage, err := FileOption(proto, "go_package")
	if err != nil {
		return nil, err
	}

	idx := strings.LastIndex(goPackage, "/")
	name := strings.Title(goPackage[idx+1:])
	module := goPackage[0:idx]

	protoIR := &ProtoGolang{
		Module:   module,
		Name:     name,
		Enums:    enums,
		Services: services,
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
