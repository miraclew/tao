package kotlin

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type typeMapper int

func (m typeMapper) Map(t *proto3.Type) (string, error) {
	if t.Scalar != proto3.None {
		return scalarToKotlinMap[t.Scalar.GoString()], nil
	} else if t.Reference != "" {
		if t.Reference == "Time" {
			return "Date", nil
		}
		if t.Reference == "Any" {
			return "Any", nil
		}
		return t.Reference, nil
	} else if t.Map != nil {
		k, _ := m.Map(t.Map.Key)
		v, _ := m.Map(t.Map.Value)
		return fmt.Sprintf("Map<%s, %s>", k, v), nil
	} else {
		return "", nil
	}
}

var scalarToKotlinMap = map[string]string{
	"Double":   "Double",
	"Float":    "Float",
	"Int32":    "Int",
	"Int64":    "Int",
	"Uint32":   "Int",
	"Uint64":   "Int",
	"Sint32":   "Int",
	"Sint64":   "Int",
	"Fixed32":  "Int",
	"Fixed64":  "Int",
	"SFixed32": "Int",
	"SFixed64": "Int",
	"Bool":     "Boolean",
	"String":   "String",
	"Bytes":    "ByteArray",
}
