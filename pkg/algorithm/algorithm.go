package algorithm

/**
These algorithms were taken from STL algorithm library to make 
a certain port on Golang
**/

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