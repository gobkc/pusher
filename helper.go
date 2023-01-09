package pusher

func Push(topic string, data any) {
	NewPusher().Push(topic, data)
}

func Subs(topic string, f func(cb SubsCallback)) {
	NewPusher().getSubscribers().Subs(topic, f)
}
