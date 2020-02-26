{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}svc

import (
	"context"
	"database/sql"
	"{{.Module}}/locator"
	"{{.Module}}/{{.Pkg}}"
	"github.com/miraclew/tao/pkg/ac"
	"github.com/miraclew/tao/pkg/ce"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type DefaultService struct {
	repo Repo
}

func NewService() *DefaultService {
	conf := locator.Config()
	db, err := sqlx.Connect("mysql", conf.MysqlAddr)
	if err != nil {
		panic(err)
	}

	s := &DefaultService{repo: &MysqlRepo{DB: db}}
	locator.Register(s.Name()+"Service", s)
	locator.Register(s.Name()+"Event", &{{.Pkg}}.EventSubscriber{Subscriber: locator.Subscriber()})

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

{{- range .Service.Methods}}
func (s *DefaultService) {{.Name}}(ctx context.Context, req *{{$.Pkg}}.{{.Request}}) (*{{$.Pkg}}.{{.Response}}, error) {
	panic("not implemented")
}
{{- end}}