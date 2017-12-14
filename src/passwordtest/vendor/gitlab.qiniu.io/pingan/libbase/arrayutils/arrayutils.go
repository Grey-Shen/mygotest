package arrayutils

import "reflect"

func Contains(s interface{}, elem interface{}) bool {
	arr := reflect.ValueOf(s)

	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
		for i := 0; i < arr.Len(); i++ {
			if arr.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func ToInterfaces(s interface{}) []interface{} {
	arr := reflect.ValueOf(s)

	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
		interfaces := make([]interface{}, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			interfaces[i] = arr.Index(i).Interface()
		}
		return interfaces
	}
	return nil
}
