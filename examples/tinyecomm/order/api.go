package order

import (
	"context"
	"github.com/miraclew/tao/pkg/component/locator"
	"github.com/miraclew/tao/pkg/pb"
	"time"
)

// Reserve import
var _ = time.Time{}
var _ = pb.Empty{}

const ServiceName = "Order"

type OrderRpc interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error)
}

func LocateOrderRpc() OrderRpc {
	return locator.Locate("OrderRpc").(OrderRpc)
}

type OrderEvent interface {
	HandleCreated(f func(ctx context.Context, req *CreatedEvent) error)
}

func LocateOrderEvent() OrderEvent {
	return locator.Locate("OrderEvent").(OrderEvent)
}

type CreateOrderRequest struct {
	Lines []*OrderItemLine
}

type CreateOrderResponse struct {
	OrderId int64
}

type OrderItemLine struct {
	ProductId int64
	Quantity int64
}

type CreatedEvent struct {
	OrderId int64
}

type Empty struct {
}
