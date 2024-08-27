package ctxkey

import (
	"reflect"
)

type (
	Context interface {
		Value(any) any
	}

	key struct {
		name         string
		serializable bool
		t            reflect.Type
	}

	Option func(*key)
)

func Key[T any](name string, serializable bool) *key {
	return &key{name, serializable, reflect.TypeFor[T]()}
}

func (k *key) Value() (val reflect.Value) {
	switch k.t.Kind() {
	case reflect.Slice:
		val = reflect.New(reflect.SliceOf(k.t.Elem())).Elem()
	case reflect.Map:
		val = reflect.New(reflect.MapOf(k.t.Key(), k.t.Elem())).Elem()
	case reflect.Ptr:
		val = reflect.New(k.t.Elem())
	default:
		val = reflect.New(k.t).Elem()
	}
	return val
}

func (k *key) Serializable() bool {
	return k.serializable
}

func (k *key) String() string {
	return k.name
}

func Serializable(s bool) Option {
	return func(k *key) {
		k.serializable = s
	}
}
