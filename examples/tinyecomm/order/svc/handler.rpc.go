package ordersvc

import (
	"github.com/miraclew/tao/examples/tinyecomm/order"
	"github.com/miraclew/tao/pkg/ac"
	"github.com/miraclew/tao/pkg/validate"
	"github.com/miraclew/tao/pkg/handler"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type rpcHandler struct {
	Service *DefaultService
}

func (h *rpcHandler) RegisterRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {
	e.POST("/v1/order/createorder", h.CreateOrder, m...)
}

func (h *rpcHandler) CreateOrder(c echo.Context) error {
	ctx := ac.FromEcho(c)
	req := new(order.CreateOrderRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := validate.Validate(req); err != nil {
		return err
	}

	res, err := h.Service.CreateOrder(ctx, req)
	if err != nil {
		return errors.Wrap(err, "handler: CreateOrder error")
	}

	return handler.JSON(c, res, err)
}
