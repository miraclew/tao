package transport

import (
	"context"
	"github.com/miraclew/tao/examples/tinyecomm/order"
	"github.com/miraclew/tao/pkg/component"
	"github.com/miraclew/tao/pkg/validate"
)

type Transporter interface {
	Bind(i interface{}) error // bind unmarshal model from request
	Respond(res interface{}, err error) error
}

type Server interface {
	RegisterServices(svc []component.Service)
	Run()
}

// Handler connect transporter with service impl
type Handler struct {
	Service order.OrderRpc
}

func (h Handler) CreateOrder(ctx *Context) error {
	req := new(order.CreateOrderRequest)
	if err := ctx.Transporter.Bind(req); err != nil {
		return err
	}
	if err := validate.Validate(req); err != nil {
		return err
	}

	res, err := h.Service.CreateOrder(ctx, req)
	return ctx.Respond(res, err)
}

type Context struct {
	context.Context
	Transporter
}
