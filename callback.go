package pusher

import (
	"errors"
	"reflect"
)

type SubscriberCallBackImpl struct {
	Data any
}

var CallBackTypeError = errors.New(`checkIsSlice:dest must be a struct pointer`)

func (s *SubscriberCallBackImpl) Bind(data any) error {
	dataType := reflect.TypeOf(data)
	elem := reflect.ValueOf(data)
	if reflect.Pointer == dataType.Kind() {
		dataType = dataType.Elem()
		elem = elem.Elem()
	}
	if dataType.Kind() != reflect.Struct {
		return CallBackTypeError
	}
	elem.Set(reflect.ValueOf(s.Data))
	return nil
}

func (s *SubscriberCallBackImpl) Set(data any) {
	s.Data = data
}
