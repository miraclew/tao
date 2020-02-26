{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/api.API*/ -}}
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

{{range .Events}}
func (e *EventSubscriber) Handle{{.}}(f func(ctx context.Context, req *{{.}}Event) error) {
    fmt.Println("event subscriber: register {{$.Name}}.{{.}}")
    _, _ = e.Subscribe("{{$.Name}}.{{.}}", func(topic string, msg []byte) error {
        var req = new({{.}}Event)
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("{{$.Name}}"), req)
    })
}
{{end}}
