package ac

import (
	"context"
	"github.com/miraclew/tao/pkg/auth"
	"github.com/miraclew/tao/pkg/slice"
	"github.com/labstack/echo/v4"
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

type xContext struct {
	UserId string
	Source string
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
	if a.identity != nil {
		if slice.StringsContains(a.identity.Roles, "admin") || a.identity.Internal != "" {
			return true
		}
	}
	return a.internalSource != "" // in process call
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

	if ctx.Value("Client") != nil {
		internal := ctx.Value("Client").(string)
		var v *auth.Identity
		if ctx.Value(UserIdContextKey) != nil {
			v = ctx.Value(UserIdContextKey).(*auth.Identity)
		} else {
			v = &auth.Identity{
				Internal: internal,
			}
		}

		return &aContext{
			Context:  ctx,
			identity: v,
		}
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
