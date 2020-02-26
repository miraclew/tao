package generator

import "github.com/miraclew/tao/tools/tao/parser/proto3"

type GenerationContext struct {
	Proto            *proto3.Proto
	Resource         string // Resource name in camel case, e.g. FooBar
	GoPkg            string
	Name             string
	OutputFilePath   string
	TemplateFilePath string
}
