package flatten

import "reflect"

func Flatten(nested interface{}) []interface{} {
	res := make([]interface{}, 0)
    flattenHelper(reflect.ValueOf(nested), &res)
    return res
}

func flattenHelper(val reflect.Value, res *[]interface{}) {
    if !val.IsValid() {
        return
    }
    if val.Kind() == reflect.Interface && !val.IsNil() {
        flattenHelper(val.Elem(), res)
        return
    }
    switch val.Kind() {
        case reflect.Slice, reflect.Array:
        	if val.Len() == 0 {
                return
            }
        	for i := 0; i < val.Len(); i++ {
                flattenHelper(val.Index(i), res)
            }
        default:
        switch val.Kind() {
        	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Map, reflect.Slice, reflect.Interface:
            	if val.IsNil() {
                	return
            	}
        	}
        *res = append(*res, val.Interface())
    }
}
