{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}

import (
    "bytes"
    "context"
    "encoding/json"
    "net/http"
)

const URL = "{{.URL}}"

type Client struct {
    t *http.Client
}

{{range .Service.Methods}}
func (c *Client) {{.Name}}(ctx context.Context, req *{{.Request}}) (*{{.Response}}, error) {
    res := new({{.Response }})
    err := c.do(ctx, "{{.Name | lower}}", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}
{{end}}

func (c *Client) do(ctx context.Context, path string, req interface{}, res interface{}) error {
    buf := new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(req)
    if err != nil {
        return err
    }
    resp, err := c.t.Post(URL+"/"+path, "application/json", buf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(res)
    if err != nil {
        return err
    }
    return nil
}
