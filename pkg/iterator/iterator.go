package iterator

/**
This iterator returns Next func after what call
you'll get element and ok status
**/
func CreateIterator(iter []interface{}) func()(interface{}, bool) {
	nextChan := make(chan interface{})
	var currentElement int

	for range nextChan {
		currentElement++
	}

	return func() (interface{}, bool) {
		if currentElement != len(iter) {
			nextChan <- 0
			return iter[currentElement], true
		} else {
			close(nextChan)
		}
		return 0, false
	}
}
