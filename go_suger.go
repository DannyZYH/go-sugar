package main

import (
	"fmt"
	"reflect"
)

func main() {
	inArrayDemo()
	return
}

func inArrayDemo() {
	normalTypeArray := []string{"int", "uint", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "", "string", "float32", "float64"}
	fmt.Println(InArray("uint", normalTypeArray))
	fmt.Println(InArray("uint1", normalTypeArray))
}

//InArray 判断单个数值是否包含在数组或切片中
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
