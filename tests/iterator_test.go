package tests

import (
	// "fmt"
	// "os"
	"testing"

	"github.com/YarikRevich/GSL/pkg/iterator"
	"github.com/franela/goblin"
)

func TestArrayIterator(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'CreateIterator' function", func() {
		g.It("Test 'CreateIterator' with slice", func() {
			testArray := []int{1, 2, 3}

			it, err := iterator.NewIterator(testArray)
			g.Assert(err).IsNil("iterator returned non nil")
			
			result := []int{}
			for it.Next(){
				result = append(result, it.Get().(int))
			}

			g.Assert(len(result)).Equal(len(testArray), "it failed and iterated over invalid number of elements")
		})
	})
}

func TestArrayIteratorWithChan(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'CreateIterator' function", func() {
		g.It("Test 'CreateIterator' with slice", func() {
			testArray := []int{1, 2, 3}

			it, err := iterator.NewIteratorWithChan(testArray)
			it.Start()
			g.Assert(err).IsNil("iterator returned non nil")
			
			result := []int{}
			for it.Next(){
				result = append(result, it.Get().(int))
			}

			g.Assert(len(result)).Equal(len(testArray), "it failed and iterated over invalid number of elements")
		})
	})
}
