package pusher

import (
	"log/slog"
	"reflect"
	"time"
)

func (p *GoPusher) Push(data any) {
	go func() {
		delay := p.settings.Interval
		if len(p.subs) == 0 {
			for {
				p.Logger().Error(`no subscriber`, slog.Duration(`waiting`, p.settings.Interval))
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

func (p *GoPusher) Subs(ds any, callback func(cdata any)) {
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
				callback(msg)
			}
		}
	}()
}

func (p *GoPusher) Logger() *slog.Logger {
	return slog.Default().WithGroup(`pusher`)
}
