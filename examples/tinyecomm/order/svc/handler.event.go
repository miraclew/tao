package ordersvc

import (
	"context"
	"github.com/miraclew/tao/examples/tinyecomm/order"
)

type eventHandler struct {
	Service *DefaultService
}

func (h eventHandler) Register() {
	order.LocateOrderEvent().HandleCreated(h.onOrderCreated)
}

func (h eventHandler) onOrderCreated(ctx context.Context, req *order.CreatedEvent) error {
	return nil
}
