package generator

import "github.com/miraclew/tao/tools/tao/parser/proto3"

type GenerationContext struct {
	Proto            *proto3.Proto
	GoPkg            string
	Name             string
	OutputFilePath   string
	TemplateFilePath string
}
