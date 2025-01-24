package reflection

import "reflect"

func Walk(x interface{}, fn func(s string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), fn)
		}
	}
}
