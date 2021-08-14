/**
This iterator operations are taken from STL library
to be ported to Golang
**/

package iterator

import (
	"sync"
	"github.com/YarikRevich/GSL/tools"
)

const (
	//Start point of iteration
	itNPos = -1
)


type Iterator struct {
	slice       []interface{}
	currElement int64
}

func (it *Iterator) Next() bool {
	if int(it.currElement) == len(it.slice)-1 {
		return false
	}
	it.currElement++
	return true
}

func (it *Iterator) Get() interface{} {
	return it.slice[it.currElement]
}

func NewIterator(iter interface{}) (*Iterator, error) {
	it := new(Iterator)
	s, err := tools.CreateInterfaceSlice(iter)
	if err != nil {
		return nil, err
	}
	it.slice = s
	it.currElement = itNPos
	return it, nil
}

/**
It's almost the same implementation of iterator
but it uses chan to send next element as result
**/
type IteratorWithChan struct {
	sync.WaitGroup

	iter  chan interface{}
	head  interface{}
	slice []interface{}
	inc   int64
}

func (it *IteratorWithChan) Start() {
	it.iter = make(chan interface{})

	go func() {
		for range it.iter {
			it.inc++
			it.head = it.slice[it.inc]
			it.Done()
		}
	}()
}

func (it *IteratorWithChan) Stop() {
	close(it.iter)
}

func (it *IteratorWithChan) Next() bool {
	if int(it.inc) == len(it.slice)-1 {
		it.Stop()
		return false
	}

	it.Add(1)
	it.iter <- nil
	it.Wait()
	return true
}

func (it *IteratorWithChan) Get() interface{} {
	return it.head
}

func NewIteratorWithChan(iter interface{}) (*IteratorWithChan, error) {
	it := new(IteratorWithChan)
	s, err := tools.CreateInterfaceSlice(iter)
	if err != nil {
		return nil, err
	}
	it.slice = s
	it.inc = itNPos
	return it, nil
}
