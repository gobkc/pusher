package pusher

import (
	"fmt"
	"reflect"
)

type subscriberCallBack struct {
	d any
}

func (s *subscriberCallBack) Bind(data any) error {
	_, err := s.checkKind(data)
	if err != nil {
		return fmt.Errorf("bind:%w", err)
	}
	reflect.ValueOf(data).Elem().Set(reflect.ValueOf(s.d))
	return nil
}

func (s *subscriberCallBack) checkKind(dest interface{}) (kind reflect.Kind, err error) {
	switch reflect.TypeOf(dest).Kind() {
	case reflect.Ptr:
		kind = reflect.ValueOf(dest).Elem().Kind()
		switch kind {
		case reflect.Slice:
		case reflect.Struct:
		case reflect.String:
		case reflect.Int:
		case reflect.Int32:
		case reflect.Int64:
		case reflect.Float32:
		case reflect.Float64:
		default:
			err = fmt.Errorf("checkIsSlice:dest must be a slice/struct/string/int/int32/int64/float32/float64 pointer")
			return
		}
	default:
		err = fmt.Errorf("checkIsSlice:dest must be a slice/struct pointer")
		return
	}
	return
}
