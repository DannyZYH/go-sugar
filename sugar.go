package sugar

import (
	"reflect"
)

// InArray
// Determines whether a single value is contained in an array or slice
// 判断单个数值是否包含在数组或切片中
func InArray(val interface{}, array interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				return
			}
		}
	}
	return
}
