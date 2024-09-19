package pusher

import "time"

type Pusher interface {
	Push(data any)
	//Subs
	//d: this parameter represents the data used by the Push function.
	Subs(d any, f func(cb Subscriber))
	Register(func(list []*Item))
}

type Subscriber interface {
	Bind(data any) error
}

type Item struct {
	Request  any
	CallBack func(cb Subscriber)
}

type Setting struct {
	Interval        time.Duration
	ConcurrentLimit int
	EnabledLog      bool
}

type GoPusher struct {
	settings *Setting
	subs     []*subscriber
}

type subscriber struct {
	msg chan any
	ds  any //data struct
}
