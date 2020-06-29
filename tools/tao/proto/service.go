package proto

import "strings"

const (
	ServiceUnknown ServiceType = iota
	ServiceRpc
	ServiceSocket
	ServiceEvent
	ServiceQueue
)

var serviceTypes = [...]string{"Rpc", "Socket", "Event", "Queue"}

type ServiceType int

func (st ServiceType) String() string {
	return serviceTypes[st]
}

func ParseServiceName(serviceName string) (st ServiceType, name string) {
	for i, v := range serviceTypes {
		idx := strings.LastIndex(serviceName, v)
		if idx <= 0 {
			continue
		}

		name = serviceName[idx:]
		st = ServiceType(i + 1)
		break
	}
	return
}
