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
