package array

import (
	"fmt"
	"reflect"
)

func Contain(arr1, arr2 interface{}) bool {
	arr1V := reflect.ValueOf(arr1)
	arr2V := reflect.ValueOf(arr2)
	switch arr1V.Type().Kind() {
	case reflect.Slice:
		switch arr2V.Type().Kind() {
		case reflect.Slice:
			if arr1V.Len() < arr2V.Len() {
				return false
			}
			tmp := make(map[interface{}]int)
			for i := 0; i < arr1V.Len(); i++ {
				v := arr1V.Index(i).Interface()
				tmp[v] += 1
			}
			for i := 0; i < arr2V.Len(); i++ {
				v := arr2V.Index(i).Interface()
				c, ok := tmp[v]
				if ok && c > 0 {
					tmp[v] -= 1
					continue
				}
				return false
			}
			return true
		default:
			for i := 0; i < arr2V.Len(); i++ {
				if reflect.DeepEqual(arr1V.Index(i).Interface(), arr2) {
					return true
				}
			}
			return false
		}
	default:
		panic(fmt.Errorf("unsupported type %v", arr1V.Type().Kind()))
	}
}

func SubFrom(arr1, arr2 interface{}) int {
	arr1V := reflect.ValueOf(arr1)
	arr2V := reflect.ValueOf(arr2)
	if arr1V.Type().Kind() != arr2V.Type().Kind() && arr1V.Type().Kind() != reflect.Slice {
		return -2
	}
	if arr1V.Len() < arr2V.Len() || arr2V.Len() == 0 {
		return -1
	}
	started := -1
	for i := 0; i < arr1V.Len(); i++ {
		if !reflect.DeepEqual(arr1V.Index(i).Interface(), arr2V.Index(0).Interface()) {
			continue
		}
		if i+arr2V.Len() > arr1V.Len() {
			continue
		}
		started = i
		for j := 1; j < arr2V.Len(); j++ {
			if !reflect.DeepEqual(arr1V.Index(started+j).Interface(), arr2V.Index(j).Interface()) {
				started = -1
				break
			}
		}
	}
	return started
}

func Equal(arr1, arr2 interface{}) bool {
	arr1V := reflect.ValueOf(arr1)
	arr2V := reflect.ValueOf(arr2)
	if arr1V.Type().Kind() != arr2V.Type().Kind() && arr1V.Type().Kind() != reflect.Slice {
		return false
	}
	if arr1V.Len() != arr2V.Len() {
		return false
	}
	tmp := make(map[interface{}]int)
	for i := 0; i < arr1V.Len(); i++ {
		v := arr1V.Index(i).Interface()
		tmp[v] += 1
	}
	for i := 0; i < arr2V.Len(); i++ {
		v := arr2V.Index(i).Interface()
		if _, ok := tmp[v]; !ok {
			return false
		}
		tmp[v] -= 1
		if tmp[v] == 0 {
			delete(tmp, v)
		}
	}
	return len(tmp) == 0
}
