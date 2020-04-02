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

func (s *DefaultService) Create(ctx context.Context, req *{{.Pkg}}.CreateRequest) (*{{.Pkg}}.CreateResponse, error) {
    c := ac.FromContext(ctx)
    req.{{.Name|title}}.UserId = c.UserID()
    id, err := s.repo.Create(ctx, req.{{.Name|title}})
    if err != nil {
        return nil, err
    }

    req.{{.Name|title}}.Id = id
    err = locator.Publisher().Publish(s.Name()+".Created", &{{.Pkg}}.CreatedEvent{Data: req.{{.Name|title}} })
    if err != nil {
        return nil, err
    }
    return &{{.Pkg}}.CreateResponse{Id: id}, nil
}

func (s *DefaultService) Delete(ctx context.Context, req *{{.Pkg}}.DeleteRequest) (*{{.Pkg}}.DeleteResponse, error) {
    res, err := s.repo.Get(ctx, &GetParams{Id: req.Id})
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ce.NewNotFound("not found")
        }
        return nil, err
    }
    c := ac.FromContext(ctx)
    if !c.Privilege() && c.UserID() != res.UserId {
        return nil, ce.NewForbidden("forbidden")
    }

    err = s.repo.Delete(ctx, req.Id)
    if err != nil {
        return nil, err
    }

    err = locator.Publisher().Publish(s.Name()+".Deleted", {{.Pkg}}.DeletedEvent{Data: res})
    if err != nil {
        return nil, err
    }
    return &{{.Pkg}}.DeleteResponse{Result: "OK"}, nil
}

func (s *DefaultService) Update(ctx context.Context, req *{{.Pkg}}.UpdateRequest) (*{{.Pkg}}.UpdateResponse, error) {
    res, err := s.repo.Get(ctx, &GetParams{Id: req.Id})
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ce.NewNotFound("not found")
        }
        return nil, err
    }
    c := ac.FromContext(ctx)
    if !c.Privilege() && c.UserID() != res.UserId {
        return nil, ce.NewForbidden("forbidden")
    }
    err = s.repo.Update(ctx, req.Values, req.Id)
    if err != nil {
        return nil, err
    }

    err = locator.Publisher().Publish(s.Name()+".Updated", {{.Pkg}}.UpdatedEvent{Id: req.Id, Values: req.Values})
    if err != nil {
        return nil, err
    }
    return &{{.Pkg}}.UpdateResponse{ Result: "OK" }, nil
}

func (s *DefaultService) Query(ctx context.Context, req *{{.Pkg}}.QueryRequest) (*{{.Pkg}}.QueryResponse, error) {
    if req.Limit == 0 {
        req.Limit = 10
    }

    res, err := s.repo.Query(ctx, &QueryParams{
        Filter: req.Filter,
        Sort:   req.Sort,
        Offset: req.Offset,
        Limit:  req.Limit,
    })
    if err != nil {
        return nil, err
    }

    return &{{.Pkg}}.QueryResponse{ Results: res }, nil
}

func (s *DefaultService) Get(ctx context.Context, req *{{.Pkg}}.GetRequest) (*{{.Pkg}}.GetResponse, error) {
    res, err := s.repo.Get(ctx, &GetParams{Id: req.Id, Filter: req.Filter})
    if err != nil {
        if err == ErrRecordNotFound {
            return nil, ce.NewNotFound("record not found")
        }
        return nil, err
    }

    return &{{.Pkg}}.GetResponse{Result: res}, nil
}
