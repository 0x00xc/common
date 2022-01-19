package array

import (
	"errors"
	"reflect"
)

// In
// 判断数组中是否包含指定元素，数组中元素类型和指定值的类型必须匹配；
// 支持 slice, array, string, map；
// list 类型为 string 时，v的类型必须为 byte 或 uint8。
func In(list interface{}, v interface{}) bool {
	val := reflect.ValueOf(list)
	switch val.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if item.Interface() == v {
				return true
			}
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			if iter.Value().Interface() == v {
				return true
			}
		}
	default:
		panic("not array")
	}
	return false
}

func SafeIn(list, v interface{}) (bool, error) {
	val := reflect.ValueOf(list)
	switch val.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if item.Interface() == v {
				return true, nil
			}
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			if iter.Value().Interface() == v {
				return true, nil
			}
		}
	default:
		return false, errors.New("not array")
	}
	return false, nil
}

func InInt64(array []int64, v int64) bool {
	for _, a := range array {
		if a == v {
			return true
		}
	}
	return false
}
