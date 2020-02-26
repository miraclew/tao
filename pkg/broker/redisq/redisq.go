package redisq

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/miraclew/tao/pkg/broker"
)

func New(rc *redis.Client) (broker.MessageBroker, error) {
	return &redisPubSub{rc: rc}, nil
}

type redisPubSub struct {
	rc *redis.Client
}

func (r redisPubSub) Publish(topic string, msg interface{}) error {
	v, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return r.rc.Publish(topic, v).Err()
}

func (r redisPubSub) Subscribe(topic string, hf broker.HandleFunc) (int, error) {
	ch := r.rc.Subscribe(topic).Channel()
	for message := range ch {
		_ = hf(topic, []byte(message.Payload))
	}

	return 0, nil
}

func (r redisPubSub) Unsubscribe(topic string, id int) error {
	return nil
}

func (r redisPubSub) Close() error {
	return r.rc.Close()
}
