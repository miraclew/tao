package memq

import (
	"encoding/json"
	"fmt"

	"github.com/miraclew/tao/pkg/broker"
)

type memQ struct {
	handlers map[string]HFS
	counter  int
}

type handler struct {
	hf broker.HandleFunc
	id int
}

type HFS []handler

func New() broker.MessageBroker {
	return &memQ{handlers: make(map[string]HFS)}
}

func (m memQ) Publish(topic string, msg interface{}) error {
	v, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	hs, ok := m.handlers[topic]
	if !ok {
		return nil
	}
	for _, h := range hs {
		err := h.hf(topic, v)
		if err != nil {
			fmt.Printf("memq: call hf err %s\n", err)
		}
	}
	return nil
}

func (m memQ) Subscribe(topic string, hf broker.HandleFunc) (int, error) {
	hs, ok := m.handlers[topic]
	if !ok {
		hs = make(HFS, 0)
	}

	m.counter++

	hs = append(hs, handler{
		hf: hf,
		id: m.counter,
	})
	m.handlers[topic] = hs
	return m.counter, nil
}

func (m memQ) Unsubscribe(topic string, id int) error {
	hs, ok := m.handlers[topic]
	if !ok {
		return nil
	}

	idx := -1
	for i, h := range hs {
		if h.id == id {
			idx = i
		}
	}
	if idx > -1 {
		hs = append(hs[0:idx], hs[idx+1:]...)
	}
	m.handlers[topic] = hs
	return nil
}

func (m memQ) Close() error {
	return nil
}
