package pusher

import (
	"log/slog"
	"reflect"
	"time"
)

func NewPusher(functions ...func(settings *Setting)) Pusher {
	settings := &Setting{
		Interval:        2 * time.Second,
		ConcurrentLimit: 2000,
		EnabledLog:      true,
	}
	for _, f := range functions {
		f(settings)
	}
	return &GoPusher{
		settings: settings,
	}
}

func (p *GoPusher) Push(data any) {
	go func() {
		delay := p.settings.Interval
		if len(p.subs) == 0 {
			for {
				slog.Default().Error(`no subscriber`, slog.Duration(`waiting`, p.settings.Interval))
				if len(p.subs) > 0 {
					p.Push(data)
					return
				}
				time.Sleep(delay)
			}
		}
		for _, sub := range p.subs {
			if isSameStructType(data, sub.ds) {
				sub.msg <- data
			}
		}
	}()
}

func (p *GoPusher) Subs(ds any, f func(cb Subscriber)) {
	if reflect.TypeOf(ds).Kind() == reflect.Pointer {
		ds = reflect.ValueOf(ds).Elem().Interface()
	}
	ns := subscriber{
		msg: make(chan interface{}, p.settings.ConcurrentLimit),
		ds:  ds,
	}
	p.subs = append(p.subs, &ns)
	go func() {
		for {
			select {
			case msg := <-ns.msg:
				cb := &SubscriberCallBackImpl{Data: msg}
				f(cb)
			}
		}
	}()
}

func (p *GoPusher) Register(f func(list []*Item)) {
	var list []*Item
	f(list)
	for _, item := range list {
		p.Subs(item.Request, item.CallBack)
	}
}
