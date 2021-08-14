/**
These algorithm operations are taken from STL library to 
be ported to Golang
**/

package algorithm


func ForEachOneD(iter []interface{}, callback func(index int, element interface{})){
	for index, element := range iter{
		callback(index, element)
	}
}

func ForEachTwoD(iter [][]interface{}, callback func(index int, element []interface{})){
	for index, element := range iter{
		callback(index, element)
	}
}