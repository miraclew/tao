package ir

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"fmt"
)

func FileOption(proto *proto3.Proto, option string) (string, error) {
	for _, entry := range proto.Entries {
		if entry.Option != nil && entry.Option.Name == option {
			return *entry.Option.Value.String, nil
		}
	}
	return "", fmt.Errorf("option %s undefined", option)
}
