{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"{{.Module}}/{{.Pkg}}"
)

type eventHandler struct {
	Service *DefaultService
}

func (h eventHandler) Register() {
	// example:
	{{.Pkg}}.Locate{{.Name}}Event().HandleCreated(h.on{{.Name}}Created)
}

// example:
func (h eventHandler) on{{.Name}}Created(ctx context.Context, req *{{.Pkg}}.CreatedEvent) error {
	return nil
}
