package openapiv3

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

func MapField(field *proto3.Field) Schema {
	s := Schema{
		Name:       field.Name,
		Ref:        "",
		Type:       "",
		Items:      nil, // type: array
		Properties: nil, // type: object
	}
	if field.Repeated {
		s.Type = "array"
		s.Items = &Schema{ // support ref only
			Ref: MapType(field.Type),
		}
	} else if field.Type.Reference != "" {
		if field.Type.Reference == "Time" {
			s.Type = "string"
		} else {
			s.Ref = MapType(field.Type)
		}
	} else {
		s.Type = MapType(field.Type)
	}

	return s
}

func MapType(t *proto3.Type) string {
	if t.Scalar != proto3.None {
		return scalarToGolangMap[t.Scalar.GoString()]
	} else if t.Reference != "" {
		return "#/components/schemas/" + t.Reference
	} else if t.Map != nil {
		return "object"
	} else {
		return ""
	}
}

var scalarToGolangMap = map[string]string{
	"Double":   "number",
	"Float":    "number",
	"Int32":    "integer",
	"Int64":    "integer",
	"Uint32":   "integer",
	"Uint64":   "integer",
	"Sint32":   "integer",
	"Sint64":   "integer",
	"Fixed32":  "integer",
	"Fixed64":  "integer",
	"SFixed32": "integer",
	"SFixed64": "integer",
	"Bool":     "boolean",
	"String":   "string",
	"Bytes":    "bytes",
}
