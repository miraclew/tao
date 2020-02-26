package golang

import (
	"fmt"

	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type TypeMapper int

func (m TypeMapper) Map(t *proto3.Type) (string, error) {
	if t.Scalar != proto3.None {
		return scalarToGolangMap[t.Scalar.GoString()], nil
	} else if t.Reference != "" {
		if t.Reference == "Time" {
			return "time.Time", nil
		}
		if t.Reference == "Any" {
			return "interface{}", nil
		}
		return t.Reference, nil
	} else if t.Map != nil {
		k, _ := m.Map(t.Map.Key)
		v, _ := m.Map(t.Map.Value)
		return fmt.Sprintf("map[%s]%s", k, v), nil
	} else {
		return "", nil
	}
}

var scalarToGolangMap = map[string]string{
	"Double":   "float64",
	"Float":    "float32",
	"Int32":    "int32",
	"Int64":    "int64",
	"Uint32":   "uint32",
	"Uint64":   "uint64",
	"Sint32":   "int32",
	"Sint64":   "int64",
	"Fixed32":  "int32",
	"Fixed64":  "int64",
	"SFixed32": "int32",
	"SFixed64": "int64",
	"Bool":     "bool",
	"String":   "string",
	"Bytes":    "[]byte",
}
