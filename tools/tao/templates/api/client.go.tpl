{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}

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
    return "{{.Name}}Service"
}

{{range .Service.Methods -}}
func (c *Client) {{.Name}}(ctx context.Context, req *{{.Request}}) (*{{.Response}}, error) {
    res := new({{.Response }})
    err := c.do(ctx, "/v1/{{$.Name|lower}}/{{.Name | lower}}", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}
{{end}}

func (c *Client) do(ctx context.Context, path string, req interface{}, res interface{}) error {
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
