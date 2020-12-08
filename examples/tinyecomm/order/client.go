package order

import (
    "bytes"
    "context"
    "fmt"
    "github.com/miraclew/tao/pkg/ce"
    "github.com/miraclew/tao/pkg/pb"
    "github.com/miraclew/tao/pkg/ac"
    "github.com/miraclew/tao/pkg/auth"
    "github.com/miraclew/tao/pkg/broker"
    "encoding/json"
    "github.com/pkg/errors"
    "io/ioutil"
    "net/http"
    "time"
)

var _ = pb.Empty{}
var _ = fmt.Sprintf
var _ = broker.PublisherComponent


type OrderRpcClient struct {
    http.Client
    BaseUrl string
    Issuer auth.Issuer
}

func (c *OrderRpcClient) Name() string {
    return "OrderRpcService"
}

func (c *OrderRpcClient) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
    res := new(CreateOrderResponse)
    err := c.do(ctx, "/v1/order/createorder", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}



func (c *OrderRpcClient) do(ctx context.Context, path string, req interface{}, res interface{}) error {
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

    httpReq, err := http.NewRequest("POST", c.BaseUrl + path, buf)
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

type OrderEventClient struct {
    broker.Subscriber
}

func (e *OrderEventClient) Name() string {
    return "OrderEventEvent"
}

func (e *OrderEventClient) HandleCreated(f func(ctx context.Context, req *CreatedEvent) error) {
    fmt.Println("event client: subscribe OrderEvent.CreatedEvent")
    _, _ = e.Subscribe("OrderEvent.CreatedEvent", func(topic string, msg []byte) error {
        var req = new(CreatedEvent)
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("OrderEvent"), req)
    })
}


