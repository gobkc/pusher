package pusher

func Push(topic string, data any) {
	NewPusher().Push(topic, data)
}

func Subs(topic string, f func(cb SubsCallback)) {
	NewPusher().getSubscribers().Subs(topic, f)
}

type Item struct {
	topic string
	fc    func(cb SubsCallback)
}

func Register(subs ...Item) {
	for _, sub := range subs {
		NewPusher().getSubscribers().Subs(sub.topic, sub.fc)
	}
}
