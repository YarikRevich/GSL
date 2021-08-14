package tools

import (
	"reflect"
	"errors"
)


func CreateInterfaceSlice(iter interface{}) ([]interface{}, error) {
	var r []interface{}
	value := reflect.ValueOf(iter)
	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return nil, errors.New("iter should be only slice and array")
	}
	for i := 0; i < value.Len(); i++ {
		r = append(r, value.Index(i).Interface())
	}
	return r, nil
}