package redisq

import (
	"github.com/miraclew/tao/pkg/broker"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

func New(rc *redis.Client, env string) (broker.MessageBroker, error) {
	return &redisPubSub{rc: rc, env: env}, nil
}

type redisPubSub struct {
	rc  *redis.Client
	env string
}

func (r redisPubSub) Publish(topic string, msg interface{}) error {
	v, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return r.rc.Publish(fmt.Sprintf("%s:%s", r.env, topic), v).Err()
}

func (r redisPubSub) Subscribe(topic string, hf broker.HandleFunc) (int, error) {
	ch := r.rc.Subscribe(fmt.Sprintf("%s:%s", r.env, topic)).Channel()
	go func() {
		for message := range ch {
			err := hf(topic, []byte(message.Payload))
			if err != nil {
				log.WithError(err).WithFields(log.Fields{"topic": topic, "payload": message.Payload}).Errorf("handle event error")
			}
		}
		log.Info("redisPubSub: subscribe goroutine end")
	}()

	return 0, nil
}

func (r redisPubSub) Unsubscribe(topic string, id int) error {
	return nil
}

func (r redisPubSub) Close() error {
	log.Println("redisPubSub: close")
	return r.rc.Close()
}
