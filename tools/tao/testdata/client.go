package demoservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/miraclew/tao/pkg/ac"
	"github.com/miraclew/tao/pkg/auth"
	"github.com/miraclew/tao/pkg/broker"
	"github.com/miraclew/tao/pkg/ce"
	"github.com/miraclew/tao/pkg/pb"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

var _ = pb.Empty{}

type DemoRpcClient struct {
	http.Client
	BaseUrl string
	Issuer  auth.Issuer
}

func (c *DemoRpcClient) Name() string {
	return "DemoRpcService"
}

func (c *DemoRpcClient) Create(ctx context.Context, req *NewThing) (*NewThingResult, error) {
	res := new(NewThingResult)
	err := c.do(ctx, "/v1/demoservice/create", req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DemoRpcClient) do(ctx context.Context, path string, req interface{}, res interface{}) error {
	ctx2 := ac.FromContext(ctx)

	token, _, err := c.Issuer.Issue(ctx2.Identity(), time.Minute*10)
	if err != nil {
		return errors.Wrap(err, "client: issue token error")
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", c.BaseUrl+path, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", token)
	httpReq.Header.Set("Client", ctx2.Internal())
	resp, err := c.Do(httpReq)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		s, _ := ioutil.ReadAll(resp.Body)
		return &ce.Error{
			Code:    resp.StatusCode,
			Message: string(s),
		}
	}

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

type DemoSocketClient struct {
	broker.Client
}

func (e *DemoSocketClient) Name() string {
	return "DemoSocketEvent"
}

func (e *ClientMessageClient) HandleClientMessage(f func(ctx context.Context, req *ClientMessage) error) {
	fmt.Println("event client: subscribe DemoService.ClientMessage")
	_, _ = e.Subscribe("DemoService.ClientMessage", func(topic string, msg []byte) error {
		var req = new(ClientMessage)
		err := json.Unmarshal(msg, req)
		if err != nil {
			return err
		}
		return f(ac.NewInternal("DemoService"), req)
	})
}
