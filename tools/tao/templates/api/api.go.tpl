{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}

import (
	"context"
	_ "time"
)

const ServiceName = "{{.Name}}"
{{ range .Enums -}}
type {{.Name}} int

func (v {{.Name}}) String() string {
	switch v {
	{{- $type := .Name }}
	{{- range .Values}}
	case {{$type}}{{.Name}}:
		return "{{.Name}}"
	{{- end}}
	default:
		return "Unknown"
	}
}

const (
{{- $type := .Name }}
{{- range .Values}}
{{$type}}{{.Name}} {{$type}} = {{.Value}}
{{- end}}
)
{{- end}}

type {{.Service.Name}} interface {
{{- range .Service.Methods}}
	{{.Name}}(ctx context.Context, req *{{.Request}}) (*{{.Response}}, error)
{{- end}}
}

{{if .Event}}
type {{.Event.Name}} interface {
{{- range .Event.Methods}}
	{{.Name}}(f func(ctx context.Context, req *{{.Request}}) error)
{{- end}}
}
{{end}}

{{ range .Messages -}}
{{- $m := . -}}
type {{.Name}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}}{{if $m.Model}} `db:"{{.Name}}"`{{end}}
{{- end}}
}

{{if hasSuffix "Request" .Name -}}
func (req *{{.Name}}) Validate() error {
	return nil
}
{{- end}}
{{end -}}