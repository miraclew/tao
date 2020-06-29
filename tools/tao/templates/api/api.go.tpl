{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}

import (
	"context"
	"github.com/miraclew/tao/pkg/pb"
	"time"
)

// Reserve import
var _ = time.Time{}
var _ = pb.Empty{}

const ServiceName = "{{.Name}}"
{{- range .Enums }}
type {{.Name}} int

func (v {{.Name}}) String() string {
	switch v {
	{{- $type := .Name }}
	{{- range .Values}}
	case {{$type}}{{.Name}}:
		return "{{.String}}"
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

{{- range .Services}}
{{if eq .Type 1 }}
type {{.Name}} interface {
	{{- range .Methods}}
	{{.Name}}(ctx context.Context, req *{{.Request}}) (*{{.Response}}, error)
	{{- end}}
}
{{- else if eq .Type 2 }}
type {{.Name}}Server interface {
	{{- range .Methods}}
	Handle{{.Name}}(ctx context.Context, req *{{.Request}}) error
	{{- end}}
}
{{- else if eq .Type 3 }}
type {{.Name}} interface {
	{{- range .Methods}}
	{{.Name}}(f func(ctx context.Context, req *{{.Request}}) error)
	{{- end}}
}
{{- else}}
type {{.Name}} interface {
	{{- range .Methods}}
	{{.Name}}(f func(ctx context.Context, req *{{.Request}}) error)
	{{- end}}
}
{{end}}
{{- end}}

{{- range .Messages}}
{{ $m := . }}
type {{.Name}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}}{{if .Tags}} `{{.Tags}}`{{end}}
{{- end}}
}

{{- end}}
