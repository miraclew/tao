{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/golang.ProtoGolang*/ -}}
package {{.Pkg}}

import (
    "context"
    "github.com/miraclew/tao/pkg/broker"
    "github.com/miraclew/tao/pkg/ac"
    "encoding/json"
    "fmt"
)

// Typed event subscriber proxy
type EventSubscriber struct {
    broker.Subscriber
}

func (e *EventSubscriber) Name() string {
    return "{{.Name}}Event"
}

{{if .Event -}}
{{range .Event.Methods}}
func (e *EventSubscriber) Handle{{.Name}}(f func(ctx context.Context, req *{{.Request}}) error) {
    fmt.Println("event subscriber: register {{$.Name}}.{{.Name}}")
    _, _ = e.Subscribe("{{$.Name}}.{{.Name}}", func(topic string, msg []byte) error {
        var req = new({{.Request}})
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("{{$.Name}}"), req)
    })
}
{{end -}}
{{end}}
