package test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeOf(t *testing.T) {
	i := 10
	tp := reflect.TypeOf(i)
	assert.Equal(t, reflect.Int, tp.Kind())

	value := reflect.New(tp)
	elem := value.Elem()
	elem.SetInt(20)
	assert.Equal(t, 20, elem.Interface())
}

func TestValueOf(t *testing.T) {
	i := 10
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(20)
	assert.Equal(t, 20, i)
}

func TestCreateByType(t *testing.T) {
	{
		tp := reflect.TypeFor[int]()
		ptr := reflect.New(tp)
		ptr.Elem().SetInt(100)
		assert.Equal(t, 100, ptr.Elem().Interface())
	}

	{
		tp := reflect.TypeFor[[]int]()
		value := reflect.MakeSlice(tp, 3, 3)
		value.Index(0).SetInt(0)
		value.Index(1).SetInt(1)
		value.Index(2).SetInt(2)
		assert.Equal(t, []int{0, 1, 2}, value.Interface())
	}

	{
		tp := reflect.TypeFor[map[string]int]()
		value := reflect.MakeMap(tp)
		value.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(1))
		value.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(2))
		value.SetMapIndex(reflect.ValueOf("c"), reflect.ValueOf(3))
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, value.Interface())
	}
}
