package utils

import "reflect"

func InArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}

func ArrayEqual(arr1 interface{}, arr2 interface{}) bool {
	if reflect.DeepEqual(arr1, arr2) {
		return true
	}
	return false
}
