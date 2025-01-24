package reflection

import "reflect"

func Walk(x interface{}, fn func(s string)) {
	val := getValue(x)

	numVals := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.Slice:
		numVals = val.Len()
		getField = val.Index
	case reflect.Struct:
		numVals = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	}
	for i := 0; i < numVals; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
