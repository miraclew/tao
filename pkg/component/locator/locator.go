package locator

import (
	"fmt"
)

var registry = make(map[string]interface{})

func Register(serviceName string, svc interface{}) {
	fmt.Println("locator: register " + serviceName)
	registry[serviceName] = svc
}

func Locate(serviceName string) interface{} {
	return registry[serviceName]
}
