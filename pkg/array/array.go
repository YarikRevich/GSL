package array

import (
	"errors"
	"reflect"
)

func InArray(iter interface{}, element interface{}) (bool, error) {
	var iterSlice []interface{}

	value := reflect.ValueOf(iter)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return false, errors.New("iter should be array or slice")
	}

	for i := 0; i < value.Len(); i++ {
		iterSlice = append(iterSlice, value.Index(i).Interface())
	}

	for _, v := range iterSlice {
		if v == element {
			return true, nil
		}
	}

	return false, nil
}

func AreArraysEqual(firstIter interface{}, secondIter interface{}) (bool, error) {
	var firstIterSlice, secondIterSlice  []interface{}

	value := reflect.ValueOf(firstIter)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return false, errors.New("first iter should be array or slice")
	}

	for i := 0; i < value.Len(); i++ {
		firstIterSlice = append(firstIterSlice, value.Index(i).Interface())
	}

	value = reflect.ValueOf(secondIter)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return false, errors.New("second iter should be array or slice")
	}

	for i := 0; i < value.Len(); i++{
		secondIterSlice = append(secondIterSlice, value.Index(i).Interface())
	}


	if len(firstIterSlice) != len(secondIterSlice) {
		return false, nil
	}
	for i := 0; i < len(firstIterSlice); i++ {
		if firstIterSlice[i] != secondIterSlice[i] {
			return false, nil
		}
	}

	return true, nil
}

func ReverseArray(iter []interface{}) []interface{} {
	r := make([]interface{}, len(iter))
	for d := len(iter) - 1; d != -1; d-- {
		r[len(iter)-d] = iter[d]
	}
	return r
}

func GetArraysInterception(iter []interface{}, inter []interface{}) ([]interface{}, error) {
	r := make([]interface{}, len(iter))
	if n := copy(iter, r); n == 0 {
		return nil, errors.New("iter can't be empty")
	}
	for i, v := range inter {
		if ok, err := InArray(r, v); !ok && err == nil {
			r[i] = v
		}
	}

	return r, nil
}
