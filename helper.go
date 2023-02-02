package pusher

import "fmt"

func Push(topic string, data any) {
	NewPusher().Push(topic, data)
}

func Subs(topic string, f func(cb SubsCallback)) {
	NewPusher().getSubscribers().Subs(topic, f)
}

type Item struct {
	Topic any
	Cb    func(cb SubsCallback)
}

func Register(subs []*Item) {
	for _, sub := range subs {
		tp := fmt.Sprintf("%v", sub.Topic)
		NewPusher().getSubscribers().Subs(tp, sub.Cb)
	}
}
