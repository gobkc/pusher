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

func Push[T any](data T) {
	New().Push(data)
}

func Subs[T Subscriber](receiver any, f func(cb T)) {
	New().Subs(receiver, func(msg any) {
		var nm = new(T)
		v := reflect.ValueOf(nm).Elem()
		childField := v.FieldByName("SubscriberCallBackImpl")
		var field reflect.Value
		if childField.IsValid() {
			field = childField.FieldByName("Data")
		}

		if !field.IsValid() || !field.CanSet() {
			field = v.FieldByName("Data")
		}
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(msg))
		}
		f(*nm)
	})
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

func getStructName(d any) string {
	de := reflect.ValueOf(d)
	if de.Kind() == reflect.Pointer {
		de = de.Elem()
	}
	n := de.Type().Name()
	return n
}
