package reflection

import "reflect"

func Walk(data interface{}, handler func(input string)) {
	val := GetValue(data)

	switch val.Kind() {
	case reflect.String:
		handler(val.String())
	case reflect.Slice, reflect.Array:
		WalkIndexValue(val.Len(), val.Index, handler)
	case reflect.Struct:
		WalkIndexValue(val.NumField(), val.Field, handler)
	case reflect.Map:
		WalkMapIndexValue(val, handler)
		return
	case reflect.Chan:
		WalkChanValue(val, handler)
	case reflect.Func:
		WalkFuncValue(val, handler)
	}
}

func WalkChanValue(val reflect.Value, handler func(input string)) {
	for v, ok := val.Recv(); ok; v, ok = val.Recv() {
		Walk(v.Interface(), handler)
	}
}

func WalkIndexValue(count int, getField func(int) reflect.Value, handler func(input string)) {
	for i := 0; i < count; i++ {
		Walk(getField(i).Interface(), handler)
	}
}

func WalkMapIndexValue(mapValue reflect.Value, handler func(input string)) {
	for _, k := range mapValue.MapKeys() {
		Walk(mapValue.MapIndex(k).Interface(), handler)
	}
}

func WalkFuncValue(value reflect.Value, handler func(input string)) {
	fnResult := value.Call(nil)
	for _, v := range fnResult {
		Walk(v.Interface(), handler)
	}
}

func GetValue(data interface{}) reflect.Value {
	val := reflect.ValueOf(data)

	// Ptr - Pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
