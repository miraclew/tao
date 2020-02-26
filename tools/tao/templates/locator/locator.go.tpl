{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper.Locator*/ -}}
package locator

import (
	"{{.Module}}/config"
	{{- range .Resources}}
	"{{$.Module}}/{{.Name}}"
	{{- end}}
	"github.com/miraclew/tao/pkg/broker"
	"fmt"
)

const PublisherName = "EventPublisher"
const SubscriberName = "EventSubscriber"

var registry = make(map[string]interface{})

func Register(serviceName string, svc interface{}) {
	fmt.Println("locator: register " + serviceName)
	registry[serviceName] = svc
}

func Publisher() broker.Publisher {
	return registry[PublisherName].(broker.Publisher)
}

func Subscriber() broker.Subscriber {
	return registry[SubscriberName].(broker.Subscriber)
}

func RegisterConfig(cfg *config.Config) {
	registry["Config"] = cfg
}

func Config() *config.Config {
	return registry["Config"].(*config.Config)
}

{{range .Resources}}
func {{.Name|title}}() {{.Name}}.Service {
	v, ok := registry["{{.Name|title}}Service"]
	if !ok {
		panic("{{.Name|title}}Service not register")
	}
	return v.({{.Name}}.Service)
}

{{if .HasEvent -}}
func {{.Name|title}}Event() {{.Name}}.Event {
	v, ok := registry["{{.Name|title}}Event"]
	if !ok {
		panic("{{.Name|title}}Event not register")
	}
	return v.({{.Name}}.Event)
}
{{end -}}
{{end}}