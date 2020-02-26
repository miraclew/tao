package ac

import (
	"context"

	"github.com/miraclew/tao/pkg/slice"

	"github.com/labstack/echo/v4"
	"github.com/miraclew/tao/pkg/auth"
)

const (
	UserIdContextKey = "user-identity"
)

type Context interface {
	context.Context
	UserID() int64
	Identity() *auth.Identity
	Authorization() string
	Privilege() bool
	Internal() string
}

type aContext struct {
	context.Context
	identity       *auth.Identity
	authorization  string
	internalSource string
}

func (a *aContext) Internal() string {
	return a.internalSource
}

func (a *aContext) UserID() int64 {
	return a.identity.UserID
}

func (a *aContext) Authorization() string {
	return a.authorization
}

func (a *aContext) Identity() *auth.Identity {
	return a.identity
}

func (a *aContext) Privilege() bool {
	if a.identity != nil && slice.StringsContains(a.identity.Roles, "admin") {
		return true
	}
	return a.internalSource != ""
}

func FromEcho(ctx echo.Context) context.Context {
	v := ctx.Get(UserIdContextKey)
	c := context.WithValue(ctx.Request().Context(), UserIdContextKey, v)

	return &aContext{
		Context:  c,
		identity: v.(*auth.Identity),
	}
}

func FromContext(ctx context.Context) Context {
	c, ok := ctx.(Context)
	if ok {
		return c
	}
	v := ctx.Value(UserIdContextKey).(*auth.Identity)
	return &aContext{
		Context:  ctx,
		identity: v,
	}
}

func NewInternal(source string) Context {
	return &aContext{
		Context:        context.Background(),
		internalSource: source,
	}
}
