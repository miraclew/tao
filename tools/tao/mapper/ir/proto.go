package ir

import (
	"fmt"

	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

func FileOption(proto *proto3.Proto, option string) (string, error) {
	for _, entry := range proto.Entries {
		if entry.Option != nil && entry.Option.Name == option {
			return *entry.Option.Value.String, nil
		}
	}
	return "", fmt.Errorf("option %s undefined", option)
}
