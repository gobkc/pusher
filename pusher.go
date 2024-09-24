package pusher

import (
	"github.com/gobkc/to"
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
	p.ReportEventRegistered(ds)
	go func() {
		for {
			select {
			case msg := <-ns.msg:
				p.ReportEventStart(msg)
				callback(msg)
				p.ReportEventCompleted(msg)
			}
		}
	}()
}

func (p *GoPusher) Logger() *slog.Logger {
	return slog.Default().WithGroup(`pusher`)
}

func (p *GoPusher) ReportEventStart(d any) {
	if p.settings.EnabledLog {
		event := getStructName(d)
		p.Logger().Info(`event request`, slog.String(`event`, event), slog.String(`data`, to.Json(d)))
	}
}

func (p *GoPusher) ReportEventCompleted(d any) {
	if p.settings.EnabledLog {
		event := getStructName(d)
		p.Logger().Info(`event completed`, slog.String(`event`, event))
	}
}

func (p *GoPusher) ReportEventRegistered(d any) {
	if p.settings.EnabledLog {
		event := getStructName(d)
		p.Logger().Info(`event registration`, slog.String(event, `event has been registered`))
	}
}
