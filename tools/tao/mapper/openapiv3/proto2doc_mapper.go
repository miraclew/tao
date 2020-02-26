package openapiv3

import (
	"fmt"
	"strings"

	"github.com/miraclew/tao/tools/tao/parser"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

func MapProto2Doc(proto *parser.Result) (*OpenAPIV3, error) {
	data := &OpenAPIV3{
		Info: Info{
			Version: "1.0.0",
			Title:   "DouYin service",
		},
		Paths: []Path{},
		Components: Components{
			Schemas: []Schema{},
		},
	}

	var paths []Path
	var schemas []Schema

	for _, entry := range proto.APIService.Entry {
		path := Path{
			Path: fmt.Sprintf("/v1/%ss/%s", proto.ResourceName, strings.ToLower(entry.Method.Name)),
			Methods: []Method{
				{
					Name:    "post",
					Summary: entry.Method.Name,
					Tags:    []string{proto.ResourceName},
					RequestBody: &Schema{
						Ref: entry.Method.Request.Reference,
					},
					Response: &Schema{
						Ref: entry.Method.Response.Reference,
					},
				},
			},
		}
		paths = append(paths, path)
	}
	// message -> schemas
	msgs := append([]*proto3.Message{proto.ResourceMessage}, proto.APIMessages...)
	for _, msg := range msgs {
		if parser.IsPredefinedMessage(msg.Name) || strings.HasSuffix(msg.Name, "Event") {
			continue
		}

		schema := Schema{
			Name:       msg.Name,
			Type:       "object",
			Ref:        "",
			Items:      nil,
			Properties: nil,
		}
		var properties []Schema
		for _, messageEntry := range msg.Entries {
			property := MapField(messageEntry.Field)
			properties = append(properties, property)
		}
		schema.Properties = properties

		schemas = append(schemas, schema)
	}

	data.Paths = paths
	data.Components.Schemas = schemas
	return data, nil
}
