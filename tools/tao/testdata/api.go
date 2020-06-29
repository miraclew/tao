package demoservice

import (
	"context"
	"github.com/miraclew/tao/pkg/pb"
	"time"
)

// Reserve import
var _ = time.Time{}
var _ = pb.Empty{}

const ServiceName = "DemoService"

type DemoRpc interface {
	Create(ctx context.Context, req *NewThing) (*NewThingResult, error)
}

type DemoSocketServer interface {
	HandleClientMessage(ctx context.Context, req *ClientMessage) error
}

type DemoEvent interface {
	HandleCreatedEvent(f func(ctx context.Context, req *ThingCreatedEvent) error)
}

type DemoQueue interface {
	PushMessage(f func(ctx context.Context, req *PublishMessage) error)
	PopMessage(f func(ctx context.Context, req *pb.Empty) error)
}


type ClientMessage struct {
}

type ServerMessage struct {
}

type PublishMessage struct {
}

type ThingCreatedEvent struct {
}

type NewThing struct {
	Mobile string
	Code string
}

type NewThingResult struct {
	Code string
	Version string
}
