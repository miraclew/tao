package sqlschema

import "github.com/miraclew/tao/tools/tao/parser/proto3"

func MapType(t *proto3.Type) string {
	if t.Scalar != proto3.None {
		return scalarToGolangMap[t.Scalar.GoString()]
	} else if t.Reference != "" {
		if t.Reference == "Time" {
			return "datetime"
		}
		return t.Reference
	} else {
		return ""
	}
}

var scalarToGolangMap = map[string]string{
	"Double":   "float",
	"Float":    "float",
	"Int32":    "int(11)",
	"Int64":    "int(11)",
	"Uint32":   "int(11) unsigned",
	"Uint64":   "int(11) unsigned",
	"Sint32":   "int(11)",
	"Sint64":   "int(11)",
	"Fixed32":  "int(11)",
	"Fixed64":  "int(11)",
	"SFixed32": "int(11)",
	"SFixed64": "int(11)",
	"Bool":     "boolean",
	"String":   "varchar(255)",
	"Bytes":    "binary",
}
