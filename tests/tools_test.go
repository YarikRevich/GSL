package tests

import (
	"reflect"
	"testing"
	"github.com/franela/goblin"
	"github.com/YarikRevich/GSL/tools"
)

func TestCreateInterfaceSlice(t *testing.T){
	g := goblin.Goblin(t)
	g.Describe("Testing 'CreateInterfaceSlice' function", func() {
		g.It("Test 'CreateInterfaceSlice'(slice, success)", func(){
			testArray := []int{1, 2, 3}
			slice, err := tools.CreateInterfaceSlice(testArray)
			value := reflect.ValueOf(slice)
			g.Assert(value.Kind()).Equal(reflect.SliceOf(reflect.TypeOf((*interface{})(nil))).Kind())
			g.Assert(err).IsNil("'CreateInterfaceSlice' failed and returns non-nil err")	
		})	
		g.It("Test 'CreateInterfaceSlice'(array, success)", func(){
			testArray := [...]int{1, 2, 3}
			slice, err := tools.CreateInterfaceSlice(testArray)
			value := reflect.ValueOf(slice)
			g.Assert(value.Kind()).Equal(reflect.SliceOf(reflect.TypeOf((*interface{})(nil))).Kind())
			g.Assert(err).IsNil("'CreateInterfaceSlice' failed and returns non-nil err")
		})	
	})
}