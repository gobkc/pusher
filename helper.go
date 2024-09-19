package pusher

import (
	"reflect"
)

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
