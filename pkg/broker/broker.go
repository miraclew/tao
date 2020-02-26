package broker

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
