{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"{{.Module}}/{{.Pkg}}"
	"github.com/labstack/echo/v4"
)

type DefaultService struct {
}

func NewService() *DefaultService {
	s := &DefaultService{}
	{{/*locator.Register(s.Name()+"Service", s)*/}}
	{{/*locator.Register(s.Name()+"Event", &{{.Pkg}}.EventClient{Subscriber: locator.Subscriber()})*/}}
	return s
}

func (s *DefaultService) Name() string {
	return "{{.Name}}"
}

func (s *DefaultService) RegisterEventHandler() {
	eh := eventHandler{s}
	eh.Register()
}

func (s *DefaultService) RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	h := handler{Service: s}
	h.RegisterRoutes(e, m...)
}

{{- range .Services}}
{{- range .Methods}}
func (s *DefaultService) {{.Name}}(ctx context.Context, req *{{$.Pkg}}.{{.Request}}) (*{{$.Pkg}}.{{.Response}}, error) {
	panic("not implemented")
}
{{end}}
{{- end}}
