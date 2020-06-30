{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"{{.Module}}/{{.Pkg}}"
	"github.com/miraclew/tao/pkg/ac"
	"github.com/miraclew/tao/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type handler struct {
	Service {{.Pkg}}.Service
}

func (h *handler) RegisterRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	{{- range .Service.Methods}}
	e.POST("/v1/{{$.Name|lower}}/{{.Name|lower}}", h.{{.Name}}, m...)
	{{- end}}
}

{{- range .Service.Methods}}

func (h *handler) {{.Name}}(c echo.Context) error {
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
		return errors.Wrap(err, "handler: get error")
	}

	return c.JSON(200, res)
}
{{- end}}
