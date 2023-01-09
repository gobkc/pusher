package pusher

type pushPolicy interface {
	Push(topic string, data any)
	getSubscribers() subsPolicy
}

type subsPolicy interface {
	Subs(topic string, f func(cb SubsCallback))
}

type SubsCallback interface {
	Bind(data any) error
}
