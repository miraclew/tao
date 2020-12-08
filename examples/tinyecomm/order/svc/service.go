package ordersvc

import (
	"context"
	"github.com/miraclew/tao/examples/tinyecomm/order"
	"github.com/labstack/echo/v4"
)

type DefaultService struct {
}

func NewService() *DefaultService {
	s := &DefaultService{}
	
	
	return s
}

func (s *DefaultService) Name() string {
	return "Order"
}

func (s *DefaultService) ClientContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "Client", s.Name())
}

func (s *DefaultService) RegisterEventHandler() {
	eh := eventHandler{s}
	eh.Register()
}

func (s *DefaultService) RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	h := rpcHandler{Service: s}
	h.RegisterRoutes(e, m...)
}
func (s *DefaultService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	panic("not implemented")
}

func (s *DefaultService) HandleCreated(ctx context.Context, req *order.CreatedEvent) (*order.Empty, error) {
	panic("not implemented")
}

