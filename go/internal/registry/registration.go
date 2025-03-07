package registry

import (
	"fmt"
	"reflect"
)

// When the key matches the key of a registered event, use this function.
func Register(reg Registry, v Registrable, s Serializer, d Deserializer, os []BuildOption) error {
	var key string

	t := reflect.TypeOf(v)

	switch {
	// accept: (*T)(nil)
	case t.Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil():
		key = reflect.New(t).Interface().(Registrable).Key()
	// accept: *T{}, T{}
	default:
		key = v.Key()
	}

	return RegisterKey(reg, key, v, s, d, os)
}

// When the key does not match the key of a registered event, use this function.
func RegisterKey(reg Registry, key string, v any, s Serializer, d Deserializer, os []BuildOption) error {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return reg.register(key, func() any {
		return reflect.New(t).Interface()
	}, s, d, os)
}

func RegisterFactory(reg Registry, key string, fn func() any, s Serializer, d Deserializer, os []BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("factory for item `%s` returns a nil value", key)
	}

	if t := reflect.TypeOf(fn()); t.Kind() != reflect.Ptr {
		return fmt.Errorf("factory for item `%s` does not return a pointer receiver", key)
	}

	return reg.register(key, fn, s, d, os)
}
