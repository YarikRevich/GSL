package tests

import (
	"testing"

	"github.com/YarikRevich/GSL/pkg/array"
	"github.com/franela/goblin"
)

func TestInArray(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'InArray' function", func() {
		g.It("Test 'InArray' function (success)", func() {
			ok, err := array.InArray([]int{1, 2, 3}, 1)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'InArray' function (fail)", func() {
			ok, err := array.InArray([]int{1, 2, 3}, 10)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
	})
}

func TestAreEqualArrays(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'AreEqualArrays' function", func() {
		g.It("Test 'AreEqualArrays' function (success)", func() {
			ok, err := array.AreArraysEqual([]int{1, 2, 3}, []int{1, 2, 3})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsTrue("'InArray' failed and returned false")
		})
		g.It("Test 'AreEqualArrays' function (fail)", func() {
			ok, err := array.AreArraysEqual([]int{1, 2, 3}, []int{1, 2})
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(ok).IsFalse("'InArray' failed and returned true")
		})
	})
}

func TestReverseArray(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'ReverseArray' function", func() {
		testArray := []int{1, 2, 3}
		g.It("Test 'ReverseArray'(success)", func() {
			result, err := array.ReverseArray(testArray)
			g.Assert(err).IsNil("'ReverseArray' failed and returned non nil err")
			g.Assert(result[0]).Equal(3)
			g.Assert(result[1]).Equal(2)
			g.Assert(result[2]).Equal(1)
		})
		g.It("Test 'ReverseArray'(fail)", func() {
			_, err := array.ReverseArray(10)
			g.Assert(err).IsNotNil("'ReverseArray' failed and returned nil err")
		})
	})
}

func TestGetArraysInterception(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing 'GetArraysInterception' function", func() {
		g.It("Test 'GetArraysInterception' function(success)", func() {
			iter := []int{1, 2, 3}
			inter := []int{4}
			result, err := array.GetArraysInterception(iter, inter)
			g.Assert(err).IsNil("'InArray failed and returned non nil err'")
			g.Assert(len(result)).Equal(4)
			g.Assert(result[0]).Equal(1)
			g.Assert(result[1]).Equal(2)
			g.Assert(result[2]).Equal(3)
			g.Assert(result[3]).Equal(4)
		})
		g.It("Test 'GetArraysInterception' function(fail)", func() {
			iter := []int{}
			inter := []int{}
			_, err := array.GetArraysInterception(iter, inter)
			g.Assert(err).IsNotNil("'InArray failed and returned nil err'")
		})
	})
}
