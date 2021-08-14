/**
This array operations are taken from STL library
to be ported to Golang
**/

package array

import (
	"errors"

	"github.com/YarikRevich/GSL/tools"
)

func InArray(iter interface{}, element interface{}) (bool, error) {
	iterSlice, err := tools.CreateInterfaceSlice(iter)
	if err != nil {
		return false, err
	}

	for _, v := range iterSlice {
		if v == element {
			return true, nil
		}
	}

	return false, nil
}

func AreArraysEqual(firstIter interface{}, secondIter interface{}) (bool, error) {
	firstIterSlice, err := tools.CreateInterfaceSlice(firstIter)
	if err != nil {
		return false, err
	}

	secondIterSlice, err := tools.CreateInterfaceSlice(secondIter)
	if err != nil {
		return false, err
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

func ReverseArray(iter interface{}) ([]interface{}, error) {
	iterSlice, err := tools.CreateInterfaceSlice(iter)
	if err != nil {
		return nil, err
	}

	r := make([]interface{}, len(iterSlice))
	for d := len(iterSlice) - 1; d != -1; d-- {
		r[len(iterSlice)-d-1] = iterSlice[d]
	}
	return r, nil
}

func GetArraysInterception(iter interface{}, inter interface{}) ([]interface{}, error) {
	iterSlice, err := tools.CreateInterfaceSlice(iter)
	if err != nil {
		return nil, err
	}

	interSlice, err := tools.CreateInterfaceSlice(inter)
	if err != nil {
		return nil, err
	}

	r := make([]interface{}, len(iterSlice) + len(interSlice))
	if n := copy(r, iterSlice); n == 0 {
		return nil, errors.New("iter can't be empty")
	}

	for i, v := range interSlice {
		if ok, err := InArray(r, v); !ok && err == nil {
			r[len(iterSlice)+i] = v 
		}
	}	
	return r, nil
}
