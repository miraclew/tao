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

type ClientMessage struct {
	Id int64
	UserId int64
	Type int32
	SubType string
}

type ServerMessage struct {
}

type NewThing struct {
	Mobile string
	Code string
}

type NewThingResult struct {
	Code string
	Version string
}
