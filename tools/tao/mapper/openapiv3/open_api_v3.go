package openapiv3

type OpenAPIV3 struct {
	Info       Info
	Paths      []Path
	Components Components
}

type Info struct {
	Version string
	Title   string
}

type Components struct {
	Schemas []Schema
}

type Path struct {
	Path    string
	Methods []Method
}

type Method struct {
	Name        string
	Summary     string
	Tags        []string
	RequestBody *Schema
	Response    *Schema
}

type Schema struct {
	Name       string
	Type       string
	Ref        string
	Items      *Schema
	Properties []Schema
}
