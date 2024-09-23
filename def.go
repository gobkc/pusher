package pusher

import (
	"log/slog"
	"time"
)

type Pusher interface {
	Push(data any)
	//Subs
	//d: this parameter represents the data used by the Push function.
	Subs(d any, callback func(msg any))
	Logger() *slog.Logger
}

type Subscriber interface {
	Bind(data any) error
	Set(data any)
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
