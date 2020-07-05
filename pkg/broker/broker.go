package broker

import "github.com/miraclew/tao/pkg/component/locator"

type Publisher interface {
	Publish(topic string, msg interface{}) error
}

type Subscriber interface {
	Subscribe(topic string, hf HandleFunc) (int, error)
	Unsubscribe(topic string, id int) error
}

type HandleFunc = func(topic string, msg []byte) error

type MessageBroker interface {
	Publisher
	Subscriber
	Close() error
}

const SubscriberComponent = "Subscriber"
const PublisherComponent = "Publisher"

func LocateSubscriber() Subscriber {
	return locator.Locate("Subscriber").(Subscriber)
}

func LocatePublisher() Publisher {
	return locator.Locate("Publisher").(Publisher)
}
