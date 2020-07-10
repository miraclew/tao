{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"{{.Module}}/{{.Pkg}}"
	"github.com/miraclew/tao/pkg/ac"
	"github.com/miraclew/tao/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

{{range .Services -}}
{{if eq .Type 1 -}}
type rpcHandler struct {
	Service *DefaultService
}

func (h *rpcHandler) RegisterRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	{{- range .Methods}}
	e.POST("/v1/{{$.Name|lower}}/{{.Name|lower}}", h.{{.Name}}, m...)
	{{- end}}
}

{{- range .Methods}}

func (h *rpcHandler) {{.Name}}(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new({{$.Pkg}}.{{.Request}})
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := validate.Validate(req); err != nil {
		return err
	}

	res, err := h.Service.{{.Name}}(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: {{.Name}} error")
	}

	return c.JSON(200, res)
}
{{- end}}
{{- end}}
{{- end}}
