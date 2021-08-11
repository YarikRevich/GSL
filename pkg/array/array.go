package array

import "errors"

func InArray(iter []interface{}, element interface{}) bool{
	for _, v := range iter{
		if v == element{
			return true
		}
	}
	return false
}

func AreArraysEqual(firstIter []interface{}, secondIter []interface{})bool{
	if len(firstIter) != len(secondIter){
		return false
	}
	for i := 0; i < len(firstIter); i++{
		if firstIter[i] != secondIter[i]{
			return false
		}
	}

	return true
}


func ReverseArray(iter []interface{})[]interface{}{
	r := make([]interface{}, len(iter))
	for d := len(iter)-1; d != -1; d--{
		r[len(iter) - d] = iter[d]
	}
	return r
}

func GetArraysInterception(iter []interface{}, inter []interface{})([]interface{}, error){
	r := make([]interface{}, len(iter))
	if n := copy(iter, r); n == 0{
		return nil, errors.New("iter can't be empty")
	}
	for i, v := range inter{
		if !InArray(r, v){
			r[i] = v
		}
	}

	return r, nil
}
