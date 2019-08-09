package utils

import "reflect"

func InArray(v interface{}, arr interface{}) (exists bool) {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return
	}

	s := reflect.ValueOf(arr)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(v, s.Index(i).Interface()) {
			return true
		}
	}
	return
}
