package tests

import (
	"testing"

	"github.com/YarikRevich/GSL/pkg/array"
	"github.com/franela/goblin"
)

func TestInArray(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'InArray' function", func() {
		g.It("Test 'InArray' function with slice (success)", func() {
			ok, err := array.InArray([]int{1, 2, 3}, 1)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'InArray' function with array (success)", func() {
			ok, err := array.InArray([...]int{1, 2, 3}, 1)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'InArray' function with slice (fail)", func() {
			ok, err := array.InArray([]int{1, 2, 3}, 10)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
		g.It("Test 'InArray' function with array (fail)", func() {
			ok, err := array.InArray([...]int{1, 2, 3}, 10)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
	})
}

func TestAreEqualArrays(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'AreEqualArrays' function", func() {
		g.It("Test 'AreEqualArrays' function with slice (success)", func() {
			ok, err := array.AreArraysEqual([]int{1, 2, 3}, []int{1, 2, 3})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'AreEqualArrays' function with array (success)", func() {
			ok, err := array.AreArraysEqual([...]int{1, 2, 3}, [...]int{1, 2, 3})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'AreEqualArrays' function with slice (fail)", func() {
			ok, err := array.AreArraysEqual([]int{1, 2, 3}, []int{1, 2})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
		g.It("Test 'AreEqualArrays' function with array (fail)", func() {
			ok, err := array.AreArraysEqual([...]int{1, 2, 3}, [...]int{1, 2})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
	})
}
