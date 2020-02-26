package dart

import (
	"fmt"

	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type typeMapper int

func (m typeMapper) Map(t *proto3.Type) (string, error) {
	if t.Scalar != proto3.None {
		return scalarToDartMap[t.Scalar.GoString()], nil
	} else if t.Reference != "" {
		if t.Reference == "Time" {
			return "DateTime", nil
		}
		if t.Reference == "Any" {
			return "dynamic", nil
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

var scalarToDartMap = map[string]string{
	"Double":   "double",
	"Float":    "double",
	"Int32":    "int",
	"Int64":    "int",
	"Uint32":   "int",
	"Uint64":   "int",
	"Sint32":   "int",
	"Sint64":   "int",
	"Fixed32":  "int",
	"Fixed64":  "int",
	"SFixed32": "int",
	"SFixed64": "int",
	"Bool":     "bool",
	"String":   "String",
	"Bytes":    "byte",
}
