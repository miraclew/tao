{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper.Locator*/ -}}
package locator

import (
	"{{.Module}}/config"
	{{- range .Resources}}
	"{{.Module}}/{{.Pkg}}"
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
func {{.Name}}() {{.Pkg}}.Service {
	v, ok := registry["{{.Name}}Service"]
	if !ok {
		panic("{{.Name}}Service not register")
	}
	return v.({{.Pkg}}.Service)
}

{{if .HasEvent -}}
func {{.Name}}Event() {{.Pkg}}.Event {
	v, ok := registry["{{.Name}}Event"]
	if !ok {
		panic("{{.Name}}Event not register")
	}
	return v.({{.Pkg}}.Event)
}
{{end -}}
{{end}}
