package pusher

import (
	"reflect"
	"sync"
	"time"
)

var once sync.Once
var ps Pusher

func New(functions ...func(settings *Setting)) Pusher {
	once.Do(func() {
		settings := &Setting{
			Interval:        2 * time.Second,
			ConcurrentLimit: 2000,
			EnabledLog:      true,
		}
		for _, f := range functions {
			f(settings)
		}
		ps = &GoPusher{
			settings: settings,
		}
	})
	return ps
}

func Push[T comparable](data T) {
	New().Push(data)
}

func Subs[T Subscriber](receiver any, f func(cb T)) {
	New().Subs(receiver, func(msg any) {
		nm := msg.(T) // new message
		nm.Set(msg)
		f(nm)
	})
}

func Register[T Subscriber](f func(list []*struct {
	Request  any
	CallBack func(cb T)
})) {
	var list []*struct {
		Request  any
		CallBack func(cb T)
	}
	f(list)
	for _, item := range list {
		Subs(item.Request, item.CallBack)
	}
}

func isSameStructType(struct1, struct2 any) bool {
	k1 := reflect.TypeOf(struct1)
	if k1.Kind() == reflect.Pointer {
		k1 = k1.Elem()
	}
	k2 := reflect.TypeOf(struct2)
	if k2.Kind() == reflect.Pointer {
		k2 = k2.Elem()
	}
	if k1.Kind() != reflect.Struct ||
		k2.Kind() != reflect.Struct {
		return false
	}
	return k1 == k2
}
