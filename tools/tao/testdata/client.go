package demoservice

import (
    "bytes"
    "context"
    "github.com/miraclew/tao/pkg/ce"
    "github.com/miraclew/tao/pkg/pb"
    "github.com/miraclew/tao/pkg/ac"
    "github.com/miraclew/tao/pkg/auth"
    "encoding/json"
    "github.com/pkg/errors"
    "io/ioutil"
    "net/http"
    "time"
)

var _ = pb.Empty{}

type Client struct {
    http.Client
    BaseUrl string
    Issuer auth.Issuer
}

func (c *Client) Name() string {
    return "DemoServiceService"
}

